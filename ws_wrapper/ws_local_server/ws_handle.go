package ws_local_server

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"open_im_sdk/pkg/log"
	utils2 "open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"open_im_sdk/ws_wrapper/utils"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type EventData struct {
	Event       string `json:"event"`
	ErrCode     int32  `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
	Data        string `json:"data"`
	OperationID string `json:"operationID"`
}

type BaseSuccessFailed struct {
	funcName    string //e.g open_im_sdk/open_im_sdk.Login
	operationID string
	uid         string
}

// 获取函数名，如Login
// e.g open_im_sdk/open_im_sdk.Login ->Login
func cleanUpfuncName(funcName string) string {
	end := strings.LastIndex(funcName, ".")
	if end == -1 {
		log.Info("", "funcName not include.", funcName)
		return ""
	}
	return funcName[end+1:]
}

// 发送错误信息到本地channel
func (b *BaseSuccessFailed) OnError(errCode int32, errMsg string) {
	log.Info("", "!!!!!!!OnError ", b.uid, b.operationID, b.funcName)
	SendOneUserMessage(EventData{cleanUpfuncName(b.funcName), errCode, errMsg, "", b.operationID}, b.uid)
}

// 发送成功信息到本地channel
func (b *BaseSuccessFailed) OnSuccess(data string) {
	log.Info("", "!!!!!!!OnSuccess ", b.uid, b.operationID, b.funcName)
	SendOneUserMessage(EventData{cleanUpfuncName(b.funcName), 0, "", data, b.operationID}, b.uid)
}

// 获取正在执行函数的名称
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

//uid->funcname->func

type WsFuncRouter struct {
	uId string
	//conn *UserConn
}

func DelUserRouter(uid string) {
	log.Info("", "DelUserRouter ", uid)
	// 如: ' '+''ios'
	sub := " " + utils.PlatformIDToName(sdk_struct.SvrConf.Platform)
	idx := strings.LastIndex(uid, sub)
	if idx == -1 {
		log.Info("", "err uid, not Web", uid, sub)
		return
	}

	// 获取没有platform标记的uid, 如‘123 IOS’ -> '123'
	uid = uid[:idx]

	UserRouteRwLock.Lock()
	defer UserRouteRwLock.Unlock()
	// UserRouteMap存放的是uid -> func name -> func
	urm, ok := UserRouteMap[uid]
	operationID := utils2.OperationIDGenerator()
	if ok {

		log.Info(operationID, "DelUserRouter logout, UnInitSDK ", uid, operationID)

		// 退出登录
		urm.wsRouter.LogoutNoCallback(uid, operationID)
		urm.wsRouter.UnInitSDK()
	} else {
		log.Info(operationID, "no found UserRouteMap: ", uid)
	}
	log.Info(operationID, "DelUserRouter delete ", uid)
	t, ok := UserRouteMap[uid]
	if ok {
		t.refName = make(map[string]reflect.Value)
	}

	// 删除映射关系
	delete(UserRouteMap, uid)
}

func GenUserRouterNoLock(uid string, batchMsg int, operationID string) *RefRouter {
	_, ok := UserRouteMap[uid]
	if ok {
		return nil
	}
	// 函数名 -> 函数方法
	RouteMap1 := make(map[string]reflect.Value, 0)
	var wsRouter1 WsFuncRouter
	wsRouter1.uId = uid

	vf := reflect.ValueOf(&wsRouter1)
	vft := vf.Type()

	mNum := vf.NumMethod()
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		log.Info(operationID, "index:", i, " MethodName:", mName)
		// 把WsFuncRouter相关的函数名对应的函数方法进行映射
		RouteMap1[mName] = vf.Method(i)
	}
	wsRouter1.InitSDK(ConfigSvr, operationID)
	log.Info(operationID, "SetAdvancedMsgListener() ", uid)
	wsRouter1.SetAdvancedMsgListener()
	if batchMsg == 1 {
		log.Info(operationID, "SetBatchMsgListener() ", uid)
		wsRouter1.SetBatchMsgListener()
	}
	// 配置listener: UserRouterMap = make(map[string]*login.LoginMgr, 0)
	wsRouter1.SetConversationListener()
	log.Info(operationID, "SetFriendListener() ", uid)
	wsRouter1.SetFriendListener()
	log.Info(operationID, "SetGroupListener() ", uid)
	wsRouter1.SetGroupListener()
	log.Info(operationID, "SetUserListener() ", uid)
	wsRouter1.SetUserListener()
	log.Info(operationID, "SetSignalingListener() ", uid)
	wsRouter1.SetSignalingListener()
	log.Info(operationID, "setWorkMomentsListener", uid)
	wsRouter1.SetWorkMomentsListener()
	var rr RefRouter
	rr.refName = RouteMap1
	rr.wsRouter = &wsRouter1
	// 在全局变量中设置uid -> func name -> func映射关系
	UserRouteMap[uid] = rr
	log.Info("", "insert UserRouteMap: ", uid)
	return &rr
}

// 发送到本地channel
func (wsRouter *WsFuncRouter) GlobalSendMessage(data interface{}) {
	SendOneUserMessage(data, wsRouter.uId)
}

// listener
// 发送数据到ws.ch（本地channel）
func SendOneUserMessage(data interface{}, uid string) {
	var chMsg ChanMsg
	chMsg.data, _ = json.Marshal(data)
	chMsg.uid = uid
	err := send2Ch(WS.ch, &chMsg, 2)
	if err != nil {
		log.Info("", "send2ch failed, ", err, string(chMsg.data), uid)
		return
	}
	log.Info("", "send response to web: ", string(chMsg.data))
}

func SendOneUserMessageForTest(data interface{}, uid string) {
	d, err := json.Marshal(data)
	log.Info("", "Marshal ", string(d))
	var chMsg ChanMsg
	chMsg.data = d
	chMsg.uid = uid
	err = send2ChForTest(WS.ch, chMsg, 2)
	if err != nil {
		log.Info("", "send2ch failed, ", err, string(chMsg.data), uid)
		return
	}
	log.Info("", "send response to web: ", string(chMsg.data))
}

// ws发送文本消息

func SendOneConnMessage(data interface{}, conn *UserConn) {
	bMsg, _ := json.Marshal(data)
	err := WS.writeMsg(conn, websocket.TextMessage, bMsg)
	log.Info("", "send response to web: ", string(bMsg), "userUid", WS.getUserUid(conn))
	if err != nil {
		log.Info("", "WS WriteMsg error", "", "userIP", conn.RemoteAddr().String(), "userUid", WS.getUserUid(conn), "error", err, "data", data)
	} else {
		log.Info("", "WS WriteMsg ok", "data", data, "userUid", WS.getUserUid(conn))
	}
}

func send2ChForTest(ch chan ChanMsg, value ChanMsg, timeout int64) error {
	var t ChanMsg
	t = value
	log.Info("", "test uid ", t.uid)
	return nil
}

// 发送消息到本地channel
func send2Ch(ch chan ChanMsg, value *ChanMsg, timeout int64) error {
	var flag = 0
	select {
	case ch <- *value:
		// 发送到channel成功
		flag = 1
	case <-time.After(time.Second * time.Duration(timeout)):
		// 发送超时
		flag = 2
	}
	if flag == 1 {
		return nil
	} else {
		log.Info("", "send cmd timeout, ", timeout, value)
		return errors.New("send cmd timeout")
	}
}

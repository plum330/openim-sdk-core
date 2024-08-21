package ws_local_server

import (
	"open_im_sdk/open_im_sdk"
)

// 用户相关的ws func router 方法集
func (wsRouter *WsFuncRouter) GetUsersInfo(userIDList string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, userIDList, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Full().GetUsersInfo(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, userIDList, operationID)
}

func (wsRouter *WsFuncRouter) SetSelfInfo(userInfo string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, userInfo, operationID, runFuncName(), nil) {
		return
	}
	userWorker.User().SetSelfInfo(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, userInfo, operationID)
}

func (wsRouter *WsFuncRouter) GetSelfUserInfo(input string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, input, operationID, runFuncName(), nil) {
		return
	}
	userWorker.User().GetSelfUserInfo(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, operationID)
}

type UserCallback struct {
	uid string
}

// 用户相关的回调方法
func (u *UserCallback) OnSelfInfoUpdated(userInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", userInfo, "0"}, u.uid)
}

// 设置用户相关的回调
func (wsRouter *WsFuncRouter) SetUserListener() {
	var u UserCallback
	u.uid = wsRouter.uId
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	userWorker.SetUserListener(&u)
}

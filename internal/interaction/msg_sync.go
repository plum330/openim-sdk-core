package interaction

import (
	"open_im_sdk/pkg/common"
	"open_im_sdk/pkg/constant"
	"open_im_sdk/pkg/db"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
)

type SeqPair struct {
	BeginSeq uint32
	EndSeq   uint32
}

type MsgSync struct {
	*db.DataBase
	*Ws
	loginUserID string
	// 消息同步时使用到的会话/消息/seq channel
	conversationCh     chan common.Cmd2Value
	PushMsgAndMaxSeqCh chan common.Cmd2Value

	selfMsgSync *SelfMsgSync
	//selfMsgSyncLatestModel *SelfMsgSyncLatestModel
	superGroupMsgSync *SuperGroupMsgSync
}

func (m *MsgSync) compareSeq() {
	operationID := utils.OperationIDGenerator()
	m.selfMsgSync.compareSeq(operationID)
	m.superGroupMsgSync.compareSeq(operationID)
}

func (m *MsgSync) doMaxSeq(cmd common.Cmd2Value) {
	m.selfMsgSync.doMaxSeq(cmd)
	// 超级群因为采用的是读扩散模式，所以同步逻辑是不一样的(按照conversation_id seq)
	m.superGroupMsgSync.doMaxSeq(cmd)
}

func (m *MsgSync) doPushMsg(cmd common.Cmd2Value) {
	msg := cmd.Value.(sdk_struct.CmdPushMsgToMsgSync).Msg
	switch msg.SessionType {
	case constant.SuperGroupChatType:
		m.superGroupMsgSync.doPushMsg(cmd)
	default:
		m.selfMsgSync.doPushMsg(cmd)
	}
}

// 读取PushMsgAndMaxSeqCh通道信息，进行处理
func (m *MsgSync) Work(cmd common.Cmd2Value) {
	switch cmd.Cmd {
	// 处理推送的消息
	case constant.CmdPushMsg:
		m.doPushMsg(cmd)
		// 处理max seq同步消息
	case constant.CmdMaxSeq:
		m.doMaxSeq(cmd)
	default:
		log.Error("", "cmd failed ", cmd.Cmd)
	}
}

func (m *MsgSync) GetCh() chan common.Cmd2Value {
	return m.PushMsgAndMaxSeqCh
}

// 消息同步流程

func NewMsgSync(dataBase *db.DataBase, ws *Ws, loginUserID string, ch chan common.Cmd2Value, pushMsgAndMaxSeqCh chan common.Cmd2Value, joinedSuperGroupCh chan common.Cmd2Value) *MsgSync {
	// PushMsgAndMaxSeqCh 消息同步本地channel
	p := &MsgSync{DataBase: dataBase, Ws: ws, loginUserID: loginUserID, conversationCh: ch, PushMsgAndMaxSeqCh: pushMsgAndMaxSeqCh}
	p.superGroupMsgSync = NewSuperGroupMsgSync(dataBase, ws, loginUserID, ch, joinedSuperGroupCh)
	p.selfMsgSync = NewSelfMsgSync(dataBase, ws, loginUserID, ch)
	//	p.selfMsgSync = NewSelfMsgSyncLatestModel(dataBase, ws, loginUserID, ch)
	// sdk启动时候，首先就要计算出需要同步的max seq
	p.compareSeq()
	// 监听处理消息同步的本地channel
	go common.DoListener(p)
	return p
}

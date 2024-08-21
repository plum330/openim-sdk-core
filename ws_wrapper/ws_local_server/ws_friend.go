package ws_local_server

import (
	"encoding/json"
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
)

type FriendCallback struct {
	uid string
}

// 类似回调函数，当friend相关变化时的回调方法，发送信息到本地channel
func (f *FriendCallback) OnFriendApplicationAdded(applyUserInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", applyUserInfo, "0"}, f.uid)
}
func (f *FriendCallback) OnFriendApplicationDeleted(applyUserInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", applyUserInfo, "0"}, f.uid)
}
func (f *FriendCallback) OnFriendApplicationAccepted(applyUserInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", applyUserInfo, "0"}, f.uid)
}
func (f *FriendCallback) OnFriendApplicationRejected(applyUserInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", applyUserInfo, "0"}, f.uid)
}
func (f *FriendCallback) OnFriendAdded(friendInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", friendInfo, "0"}, f.uid)
}
func (f *FriendCallback) OnFriendDeleted(friendInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", friendInfo, "0"}, f.uid)
}
func (f *FriendCallback) OnFriendInfoChanged(userInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", userInfo, "0"}, f.uid)
}
func (f *FriendCallback) OnBlackAdded(userInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", userInfo, "0"}, f.uid)
}
func (f *FriendCallback) OnBlackDeleted(friendInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", friendInfo, "0"}, f.uid)
}

// 设置friend变化时，对应的回调方法集
func (wsRouter *WsFuncRouter) SetFriendListener() {
	var fr FriendCallback
	fr.uid = wsRouter.uId
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	userWorker.SetFriendListener(&fr)
}

// ws func router friend相关方法集
// 1
func (wsRouter *WsFuncRouter) GetDesignatedFriendsInfo(userIDList string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, userIDList, operationID, runFuncName(), nil) {
		return
	}
	// 获取指定的friend
	userWorker.Friend().GetDesignatedFriendsInfo(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, userIDList, operationID)
}

// 1
func (wsRouter *WsFuncRouter) AddFriend(paramsReq string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, paramsReq, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().AddFriend(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, paramsReq, operationID)
}

// 获取添加自己为好友的申请列表
func (wsRouter *WsFuncRouter) GetRecvFriendApplicationList(input string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, input, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().GetRecvFriendApplicationList(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, operationID)
}

// 获取自己添加别人为好友的申请列表
func (wsRouter *WsFuncRouter) GetSendFriendApplicationList(input string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, input, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().GetSendFriendApplicationList(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, operationID)
}

// 1
// 接受好友申请
func (wsRouter *WsFuncRouter) AcceptFriendApplication(params string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, params, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().AcceptFriendApplication(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, params, operationID)
}

// 1
// 拒绝好友申请
func (wsRouter *WsFuncRouter) RefuseFriendApplication(params string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, params, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().RefuseFriendApplication(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, params, operationID)
}

// 1
func (wsRouter *WsFuncRouter) CheckFriend(userIDList string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, userIDList, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().CheckFriend(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, userIDList, operationID)
}

// 1
func (wsRouter *WsFuncRouter) DeleteFriend(friendUserID string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, friendUserID, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().DeleteFriend(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, friendUserID, operationID)
}

// 1
func (wsRouter *WsFuncRouter) GetFriendList(input string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, input, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().GetFriendList(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, operationID)
}

// 搜索还有 (本地)
func (wsRouter *WsFuncRouter) SearchFriends(input string, operationID string) {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(input), &m); err != nil {
		log.Info(operationID, utils.GetSelfFuncName(), "unmarshal failed", input, err.Error())
		wsRouter.GlobalSendMessage(EventData{cleanUpfuncName(runFuncName()), StatusBadParameter, "unmarshal failed", "", operationID})
		return
	}
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, input, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().SearchFriends(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId},
		input, operationID)
}

// 1
// 好友备注
func (wsRouter *WsFuncRouter) SetFriendRemark(remark string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, remark, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().SetFriendRemark(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, remark, operationID)
}

// 还有加入黑名单
func (wsRouter *WsFuncRouter) AddBlack(blackUserID string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, blackUserID, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().AddBlack(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, blackUserID, operationID)
}

func (wsRouter *WsFuncRouter) GetBlackList(input string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, input, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().GetBlackList(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, operationID)
}

// 从黑名单中移除
func (wsRouter *WsFuncRouter) RemoveBlack(removeUserID string, operationID string) {
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, removeUserID, operationID, runFuncName(), nil) {
		return
	}
	userWorker.Friend().RemoveBlack(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId}, removeUserID, operationID)
}

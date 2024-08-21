package open_im_sdk

import (
	"open_im_sdk/internal/login"
	"sync"
)

func init() {
	UserSDKRwLock.Lock()
	defer UserSDKRwLock.Unlock()
	UserRouterMap = make(map[string]*login.LoginMgr, 0)
}

var UserSDKRwLock sync.RWMutex

// 全局变量保存用户的LoginMgr信息 uid -> LoginMgr
var UserRouterMap map[string]*login.LoginMgr

var userForSDK *login.LoginMgr

// 获取用户对应的LoginMgr对象
func GetUserWorker(uid string) *login.LoginMgr {
	UserSDKRwLock.Lock()
	defer UserSDKRwLock.Unlock()
	v, ok := UserRouterMap[uid]
	if ok {
		return v
	}
	UserRouterMap[uid] = new(login.LoginMgr)

	return UserRouterMap[uid]
}

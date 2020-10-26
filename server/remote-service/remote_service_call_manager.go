package remoteservice

import (
	"time"

	"github.com/luckyweiwei/base/grmon"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/proto"
	cmap "github.com/orcaman/concurrent-map"
)

var (
	TimeOut = 30 * time.Second
)

type RemoteServiceCallManager struct {
	SysParamByAppMap cmap.ConcurrentMap
}

var remoteServiceCallManager *RemoteServiceCallManager = nil

func GetRemoteServiceCallManager() *RemoteServiceCallManager {
	return remoteServiceCallManager
}

func RemoteServiceCallManagerInit() {
	remoteServiceCallManager = &RemoteServiceCallManager{
		SysParamByAppMap: cmap.New(),
	}

	grm := grmon.GetGRMon()
	grm.GoLoop("SysParamByAppCall", SysParamByAppCall)
}

func (r *RemoteServiceCallManager) GetSysParamByApp(appID string) *proto.SysParamByAppResp {
	obj, ok := r.SysParamByAppMap.Get(appID)
	if !ok {
		// 有可能信息还没有更新到map，这里再调用一遍远程服务
		SysParamByApp(appID)
		obj, ok = r.SysParamByAppMap.Get(appID)
		if !ok {
			Log.Warnf("cannot find app in map, appID = %v", appID)
			return nil
		}
	}

	sysParam := obj.(*proto.SysParamByAppResp)

	return sysParam
}

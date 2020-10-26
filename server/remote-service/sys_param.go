package remoteservice

import (
	"time"

	"github.com/luckyweiwei/base/grmon"
	httpclient "github.com/luckyweiwei/base/http-client"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	"github.com/luckyweiwei/websocketserver/server/consul"
)

func SysParamByAppCall() {
	Log.Debug("enter...")

	// 得到所有的 app
	grm := grmon.GetGRMon()
	cm := connectionm.GetConnectionManager()
	keys := cm.AppMap.Keys()

	if len(keys) > 0 {
		for _, v := range keys {
			appID := v
			grm.Go("SysParamByApp", func() { SysParamByApp(appID) })
		}

	}

	time.Sleep(time.Minute)
}

func SysParamByApp(appID string) {
	addr := consul.GetBalanceAddr(proto.SysParamServerName)
	sysParamAddr := addr + proto.SysParamAddrSuf

	queryString := "appId=" + appID

	resp, body, errs := httpclient.New().
		Post(sysParamAddr).
		Timeout(TimeOut).
		Query(queryString).
		End()

	if errs != nil {
		Log.Error(errs)
		return
	}

	if resp.StatusCode != 200 {
		Log.Error("req status code != 200, resp = %v", resp)
		return
	}

	rm := GetRemoteServiceCallManager()

	var respData = &proto.SysParamByAppResp{}
	err := utils.DecodeFromJson(body, respData)
	if err != nil {
		Log.Error(err)
		return
	}

	rm.SysParamByAppMap.Set(appID, respData)
}

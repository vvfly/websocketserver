package remoteservice

import (
	httpclient "github.com/luckyweiwei/base/http-client"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/consul"
)

func GetUserWeekStarByApp(appID string) *proto.UserWeekStarByAppResp {
	addr := consul.GetBalanceAddr(proto.UserServerName)
	userWeekStarAddr := addr + proto.UserWeekStarAddrSuf

	queryString := "appId=" + appID

	resp, body, errs := httpclient.New().
		Get(userWeekStarAddr).
		Timeout(TimeOut).
		Query(queryString).
		End()

	if errs != nil {
		Log.Error(errs)
		return nil
	}

	if resp.StatusCode != 200 {
		Log.Error("req status code != 200, resp = %v", resp)
		return nil
	}

	var respData = &proto.UserWeekStarByAppResp{}
	err := utils.DecodeFromJson(body, respData)
	if err != nil {
		Log.Error(err)
		return nil
	}

	return respData
}

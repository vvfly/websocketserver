package remoteservice

import (
	httpclient "github.com/luckyweiwei/base/http-client"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/consul"
)

// 是否播放贵族动画
func GetPlayNobilityEnterAnimationByLive(liveID string) string {
	addr := consul.GetBalanceAddr(proto.NobilityServerName)
	playNobilityEnterAnimationAddr := addr + proto.PlayNobilityEnterAnimationAddrSuf

	queryString := "liveId=" + liveID

	resp, body, errs := httpclient.New().
		Get(playNobilityEnterAnimationAddr).
		Timeout(TimeOut).
		Query(queryString).
		End()

	if errs != nil {
		Log.Error(errs)
		return ""
	}

	if resp.StatusCode != 200 {
		Log.Error("req status code != 200, resp = %v", resp)
		return ""
	}

	return body
}

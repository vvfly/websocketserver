package schedule

import (
	"time"

	"github.com/luckyweiwei/base/grmon"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
)

func OnlineInfoStatistic() {
	var sleepSecond int

	cm := connectionm.GetConnectionManager()
	totalCount := cm.ConnectionMap.Count()
	if totalCount > 20000 {
		sleepSecond = 20
	} else if totalCount > 15000 {
		sleepSecond = 15
	} else if totalCount > 5000 {
		sleepSecond = 12
	} else if totalCount > 1000 {
		sleepSecond = 8
	} else if totalCount > 500 {
		sleepSecond = 6
	} else {
		sleepSecond = 4
	}

	cnfInterval := model.GetServerConfig().OnlineInfoStatisticInterval
	if cnfInterval != 8 {
		sleepSecond = cnfInterval
	}

	/*
		统计的时候先删掉当前节点数据，再批量插入
	*/
	grm := grmon.GetGRMon()
	grm.Go("RefreshAppMemberCount", RefreshAppMemberCount)
	grm.Go("RefreshOnlineLiveManager", RefreshOnlineLiveManager)
	grm.Go("RefreshOnlineVIP", RefreshOnlineVIP)
	grm.Go("RefreshSocketConnectCountInfo", RefreshSocketConnectCountInfo)
	grm.Go("RefreshOnlineCountStatistic", RefreshOnlineCountStatistic)
	grm.Go("RefreshLiveOnlineUser", RefreshLiveOnlineUser)

	Log.Infof("在线人数:%v, 统计间隔:%v", totalCount, sleepSecond)

	time.Sleep(time.Duration(sleepSecond) * time.Second)
}

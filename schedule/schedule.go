package schedule

import (
	"github.com/luckyweiwei/base/grmon"
	"github.com/robfig/cron/v3"
)

func ScheduleInit() {
	// groutine job
	grm := grmon.GetGRMon()
	grm.GoLoop("OnlineInfoStatistic", OnlineInfoStatistic)         // 在线信息统计
	grm.Go("RefreshUserOnlineInfoConn", RefreshUserOnlineInfoConn) // 用户连接，刷新用户信息
	// grm.Go("RefreshUserOnlineInfoDisConn", RefreshUserOnlineInfoDisConn) // 用户断开连接，刷新用户信息
	grm.GoLoop("RefreshUserOnlineInfoMsg", RefreshUserOnlineInfoMsg) // 用户发消息，刷新用户信息

	// cron job
	c := cron.New()
	c.AddFunc("@every 1m", HeartBeatCheck)
	c.AddFunc("@every 1h", StartHeapCheck)

	c.Start()
}

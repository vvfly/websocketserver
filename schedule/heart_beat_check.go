package schedule

import (
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/biz"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
)

func CheckLastPingTime(lastPingTime int64) bool {
	serverConfig := model.GetServerConfig()

	ct := utils.GetTimeOfS()
	if (ct - lastPingTime) > serverConfig.PingTimeInterval { // 心跳超时，认为用户已下线
		return true
	}
	return false
}

func HeartBeatCheck() {
	Log.Debug("enter...")

	cm := connectionm.GetConnectionManager()
	cm.ConnectionMap.IterCb(func(key string, v interface{}) {
		connection := v.(*connectionm.Connection)

		lastPingTime := connection.WsConnectInfo.Authen.UsrInfo.LastPingTime
		if CheckLastPingTime(lastPingTime) {
			sessionID := connection.WsConnectInfo.SessionsInfoStr
			Log.Debugf("client ping timeout.lastPingTime=%v, sessionID=%v", lastPingTime, sessionID)

			// 主动关闭连接通道
			conn := connection.Conn
			err := conn.Close()
			if err != nil {
				Log.Errorf("close connection error. err=%v", err)
				return
			}

			// 如果是主播，发送 leave消息
			enterType := connection.WsConnectInfo.Sid.EnterType
			if enterType == proto.EnterTypeAnchor {
				topic := proto.MT_LEAVE
				key := "timeout_leave_" + sessionID
				value := ""

				biz.TimeoutLeaveHandler(topic, key, value)
			}
		}
	})
}

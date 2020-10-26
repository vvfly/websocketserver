package broadcast

import (
	"github.com/luckyweiwei/base/grmon"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/helper"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
)

func BroadcastByAll(value string) {
	Log.Debug("enter...")

	var (
		grm          = grmon.GetGRMon()
		cm           = connectionm.GetConnectionManager()
		serverConfig = model.GetServerConfig()
		conns        = make([]*connectionm.Connection, 0)
	)

	eData, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), []byte(value))
	if err != nil {
		Log.Warn(err)
		return
	}
	Log.Infof("BroadcastByAll. broadcast data=%v", value)

	cm.ConnectionMap.IterCb(func(key string, v interface{}) {
		connection := v.(*connectionm.Connection)
		conns = append(conns, connection)
	})

	grm.Go("BroadcastByAll", func() {
		WriteDataToCh(conns, eData)
	})
}

func BroadcastBySessionID(sessionID string, value string) {
	Log.Debug("enter...")

	var (
		cm           = connectionm.GetConnectionManager()
		serverConfig = model.GetServerConfig()
		conns        = make([]*connectionm.Connection, 0)
	)

	eData, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), []byte(value))
	if err != nil {
		Log.Warn(err)
		return
	}
	Log.Infof("BroadcastBySessionID. sessionID=%v, broadcast data=%v", sessionID, value)

	connection := cm.GetConnection(sessionID)
	if connection == nil {
		Log.Warnf("cannot find conn in map, sessionID = %v", sessionID)
		return
	}
	conns = append(conns, connection)

	WriteDataToCh(conns, eData)
}

func BroadcastBySessionIDs(sessionIDs []string, value string) {
	Log.Debug("enter...")

	var (
		grm          = grmon.GetGRMon()
		cm           = connectionm.GetConnectionManager()
		serverConfig = model.GetServerConfig()
		conns        = make([]*connectionm.Connection, 0)
	)

	eData, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), []byte(value))
	if err != nil {
		Log.Warn(err)
		return
	}
	Log.Infof("BroadcastBySessionIDs. sessionIDs=%v, broadcast data=%v", sessionIDs, value)

	for _, v := range sessionIDs {
		connection := cm.GetConnection(v)
		if connection == nil {
			Log.Warnf("cannot find conn in map, sessionID = %v", v)
			continue
		}

		conns = append(conns, connection)
	}

	grm.Go("BroadcastBySessionIDs", func() {
		WriteDataToCh(conns, eData)
	})
}

func BroadcastByAppID(appID string, value string) {
	Log.Debug("enter...")

	var (
		grm          = grmon.GetGRMon()
		cm           = connectionm.GetConnectionManager()
		serverConfig = model.GetServerConfig()
		conns        = make([]*connectionm.Connection, 0)
	)

	eData, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), []byte(value))
	if err != nil {
		Log.Warn(err)
		return
	}
	Log.Infof("BroadcastByAppID. appID=%v, broadcast data=%v", appID, value)

	appConnection := cm.GetAppConnection(appID)
	if appConnection == nil {
		Log.Warnf("cannot find appID in map, appID = %v", appID)
		return
	}

	appConnection.Connections.IterCb(func(key string, v interface{}) {
		connection := v.(*connectionm.Connection)
		conns = append(conns, connection)
	})

	grm.Go("BroadcastByAppID", func() {
		WriteDataToCh(conns, eData)
	})
}

func BroadcastByAppIDs(appIDs []string, value string) {
	Log.Debug("enter...")

	var (
		grm          = grmon.GetGRMon()
		cm           = connectionm.GetConnectionManager()
		serverConfig = model.GetServerConfig()
		conns        = make([]*connectionm.Connection, 0)
	)

	eData, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), []byte(value))
	if err != nil {
		Log.Warn(err)
		return
	}
	Log.Infof("BroadcastByAppIDs. appIDs=%v, broadcast data=%v", appIDs, value)

	for _, v := range appIDs {
		appConnection := cm.GetAppConnection(v)
		if appConnection == nil {
			Log.Warnf("cannot find appID in map, appID = %v", v)
			continue
		}

		appConnection.Connections.IterCb(func(key string, v interface{}) {
			connection := v.(*connectionm.Connection)
			conns = append(conns, connection)
		})
	}

	grm.Go("BroadcastByAppIDs", func() {
		WriteDataToCh(conns, eData)
	})
}

func BroadcastByLiveID(liveID string, value string) {
	Log.Debug("enter...")

	var (
		grm          = grmon.GetGRMon()
		cm           = connectionm.GetConnectionManager()
		serverConfig = model.GetServerConfig()
		conns        = make([]*connectionm.Connection, 0)
	)

	eData, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), []byte(value))
	if err != nil {
		Log.Warn(err)
		return
	}
	Log.Infof("BroadcastByLiveID. liveID=%v, broadcast data=%v", liveID, value)

	roomConnection := cm.GetRommConnection(liveID)
	if roomConnection == nil {
		Log.Warnf("cannot find roomID in map, liveID = %v", liveID)
		return
	}

	roomConnection.Connections.IterCb(func(key string, v interface{}) {
		connection := v.(*connectionm.Connection)
		conns = append(conns, connection)
	})

	grm.Go("BroadcastByLiveID", func() {
		WriteDataToCh(conns, eData)
	})
}

// 入场消息广播，广播给直播间除自身以外所有人
func BroadcastEntry(liveID string, value string, sessionID string) {
	Log.Debug("enter...")

	var (
		grm          = grmon.GetGRMon()
		cm           = connectionm.GetConnectionManager()
		serverConfig = model.GetServerConfig()
		conns        = make([]*connectionm.Connection, 0)
	)

	eData, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), []byte(value))
	if err != nil {
		Log.Warn(err)
		return
	}
	Log.Infof("BroadcastEntry. liveID=%v, broadcast data=%v", liveID, value)

	roomConnection := cm.GetRommConnection(liveID)
	if roomConnection == nil {
		Log.Warnf("cannot find roomID in map, liveID = %v", liveID)
		return
	}

	roomConnection.Connections.IterCb(func(key string, v interface{}) {
		if key == sessionID {
			return
		}

		connection := v.(*connectionm.Connection)
		conns = append(conns, connection)
	})

	grm.Go("BroadcastEntry", func() {
		WriteDataToCh(conns, eData)
	})
}

func BroadcastByLiveIDs(liveIDs []string, value string) {
	Log.Debug("enter...")

	var (
		grm          = grmon.GetGRMon()
		cm           = connectionm.GetConnectionManager()
		serverConfig = model.GetServerConfig()
		conns        = make([]*connectionm.Connection, 0)
	)

	eData, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), []byte(value))
	if err != nil {
		Log.Warn(err)
		return
	}
	Log.Infof("BroadcastByLiveIDs. liveIDs=%v, broadcast data=%v", liveIDs, value)

	for _, v := range liveIDs {
		roomConnection := cm.GetRommConnection(v)
		if roomConnection == nil {
			continue
		}

		roomConnection.Connections.IterCb(func(key string, v interface{}) {
			connection := v.(*connectionm.Connection)
			conns = append(conns, connection)
		})
	}

	grm.Go("BroadcastByLiveIDs", func() {
		WriteDataToCh(conns, eData)
	})
}

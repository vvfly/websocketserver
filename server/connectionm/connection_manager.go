package connectionm

import (
	"sync"

	"github.com/gorilla/websocket"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/proto"
	cmap "github.com/orcaman/concurrent-map"
)

const (
	DATA_CHAN_MAX_LEN = 1024 // 每个连接通道最多缓存1024条消息
)

type Connection struct {
	ReadPackNum   uint64
	WritePackNum  uint64
	Conn          *websocket.Conn
	WsConnectInfo *proto.WsConnectInfo
	DataCh        chan []byte
}

type RoomConnection struct {
	Connections cmap.ConcurrentMap
}

type AppConnection struct {
	Connections cmap.ConcurrentMap
}

type ConnectionManager struct {
	Mutex         *sync.Mutex
	LastEnterTime int64 // 最后入场时间

	ConnectionMap cmap.ConcurrentMap // 按客户端
	LiveMap       cmap.ConcurrentMap // 按房间分组
	AppMap        cmap.ConcurrentMap // 按渠道分组
}

var connectionManager *ConnectionManager = nil

func ConnectionManagerInit() {
	connectionManager = &ConnectionManager{
		Mutex:         new(sync.Mutex),
		LastEnterTime: 1,
		ConnectionMap: cmap.New(),
		LiveMap:       cmap.New(),
		AppMap:        cmap.New(),
	}
}

func GetConnectionManager() *ConnectionManager {
	return connectionManager
}

func (c *ConnectionManager) NewConnection(con *websocket.Conn, wsConnInfo *proto.WsConnectInfo) *Connection {

	dataCh := make(chan []byte, DATA_CHAN_MAX_LEN)
	sessionID := wsConnInfo.SessionsInfoStr
	connection := &Connection{
		Conn:          con,
		WsConnectInfo: wsConnInfo,
		DataCh:        dataCh,
	}
	c.ConnectionMap.Set(sessionID, connection)

	liveID := wsConnInfo.Sid.LiveID
	roomConnObj, ok := c.LiveMap.Get(liveID)
	if !ok {
		roomCon := &RoomConnection{
			Connections: cmap.New(),
		}
		roomCon.Connections.Set(sessionID, connection)
		connectionManager.LiveMap.Set(liveID, roomCon)

	} else {
		roomConn := roomConnObj.(*RoomConnection)
		roomConn.Connections.Set(sessionID, connection)
	}

	appID := wsConnInfo.Sid.AppID
	appConnObj, ok := c.AppMap.Get(appID)
	if !ok {
		appCon := &AppConnection{
			Connections: cmap.New(),
		}
		appCon.Connections.Set(sessionID, connection)
		connectionManager.AppMap.Set(appID, appCon)

	} else {
		appConn := appConnObj.(*AppConnection)
		appConn.Connections.Set(sessionID, connection)
	}

	return connection
}

func (c *ConnectionManager) GetConnection(sessionID string) *Connection {
	connobj, ok := c.ConnectionMap.Get(sessionID)
	if !ok {
		Log.Warnf("cannot find conn in map, sessionID = %v", sessionID)
		return nil
	}

	connection := connobj.(*Connection)

	return connection
}

func (c *ConnectionManager) GetRommConnection(liveID string) *RoomConnection {
	connobj, ok := c.LiveMap.Get(liveID)
	if !ok {
		Log.Warnf("cannot find room in map, liveID = %v", liveID)
		return nil
	}

	connection := connobj.(*RoomConnection)

	return connection
}

func (c *ConnectionManager) GetAppConnection(appID string) *AppConnection {
	connobj, ok := c.AppMap.Get(appID)
	if !ok {
		Log.Warnf("cannot find app in map, appID = %v", appID)
		return nil
	}

	connection := connobj.(*AppConnection)

	return connection
}

func (c *ConnectionManager) DeleteConnection(sessionID string) {
	defer utils.CatchException()

	connection := c.GetConnection(sessionID)
	if connection == nil {
		Log.Warnf("cannot find conn in map, sessionID = %v", sessionID)
		return
	}
	c.ConnectionMap.Remove(sessionID)

	conn := connection.Conn
	dataCh := connection.DataCh

	err := conn.Close()
	if err != nil {
		Log.Errorf("close connection error. err=%v", err)
	}
	close(dataCh)
}

func (c *ConnectionManager) DeleteLiveMap(sessionID, liveID string) {

	roomConnection := c.GetRommConnection(liveID)
	if roomConnection == nil {
		Log.Warnf("cannot find roomID in map, liveID = %v", liveID)
		return
	}

	roomConnection.Connections.Remove(sessionID)

	if roomConnection.Connections.Count() <= 0 {
		c.LiveMap.Remove(liveID)
	}
}

func (c *ConnectionManager) DeleteAppMap(sessionID, appID string) {

	appConnection := c.GetAppConnection(appID)
	if appConnection == nil {
		Log.Warnf("cannot find appID in map, appID = %v", appID)
		return
	}

	appConnection.Connections.Remove(sessionID)

	if appConnection.Connections.Count() <= 0 {
		c.AppMap.Remove(appID)
	}
}

func (c *ConnectionManager) ClearConnection(connection *Connection) {
	sessionID := connection.WsConnectInfo.SessionsInfoStr
	liveID := connection.WsConnectInfo.Sid.LiveID
	appID := connection.WsConnectInfo.Sid.AppID

	c.DeleteAppMap(sessionID, appID)
	c.DeleteLiveMap(sessionID, liveID)
	c.DeleteConnection(sessionID)
}

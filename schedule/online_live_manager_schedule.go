package schedule

import (
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/model/dao"
	"github.com/luckyweiwei/websocketserver/model/do"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	"go.mongodb.org/mongo-driver/bson"
)

// 在线房管用户列表刷新
func RefreshOnlineLiveManager() {
	Log.Debug("enter...")

	var (
		cm       = connectionm.GetConnectionManager()
		serverID = model.GetServerConfig().ServerID
		docs     = []interface{}{}
	)

	cm.LiveMap.IterCb(func(key string, v interface{}) {
		UserIDSet := make([]string, 0)

		liveID := key
		roomConnection := v.(*connectionm.RoomConnection)
		if roomConnection == nil {
			Log.Warnf("cannot find roomID in map, liveID = %v", liveID)
			return
		}

		roomConnection.Connections.IterCb(func(key string, v interface{}) {
			connection := v.(*connectionm.Connection)

			if connection.WsConnectInfo.IsRoomManager() {
				UserIDSet = append(UserIDSet, connection.WsConnectInfo.Sid.UserID)
			}
		})

		if len(UserIDSet) > 0 {
			id := serverID + "_" + liveID

			doc := &do.OnlineLiveManager{
				ID:        id,
				ServerID:  serverID,
				LiveID:    liveID,
				UserIDSet: UserIDSet,
			}
			docs = append(docs, doc)
		}
	})

	// delete
	filter := bson.D{{"serverId", serverID}}
	_, err := dao.NewMgoOnlineLiveManagerDao().DeleteMany(filter)
	if err != nil {
		Log.Error(err)
		return
	}

	// insert
	if len(docs) > 0 {
		_, err = dao.NewMgoOnlineLiveManagerDao().InsertMany(docs)
		if err != nil {
			Log.Error(err)
			return
		}
	}

}

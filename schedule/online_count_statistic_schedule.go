package schedule

import (
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/model/dao"
	"github.com/luckyweiwei/websocketserver/model/do"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	"go.mongodb.org/mongo-driver/bson"
)

// OnlineCountStatistic表 刷新
func RefreshOnlineCountStatistic() {
	Log.Debug("enter...")

	var (
		cm       = connectionm.GetConnectionManager()
		serverID = model.GetServerConfig().ServerID
		docs     = []interface{}{}
	)

	cm.LiveMap.IterCb(func(key string, v interface{}) {
		var (
			totalCount          int64
			guardCount          int64
			monthGuardCount     int64
			yearGuardCount      int64
			roomManagerCount    int64
			nobilityCount       int64
			nobilityLevel1Count int64
			nobilityLevel2Count int64
			nobilityLevel3Count int64
			nobilityLevel4Count int64
			nobilityLevel5Count int64
			nobilityLevel6Count int64
			nobilityLevel7Count int64
		)

		liveID := key
		roomConnection := v.(*connectionm.RoomConnection)
		if roomConnection == nil {
			Log.Warnf("cannot find roomID in map, liveID = %v", liveID)
			return
		}

		roomConnection.Connections.IterCb(func(key string, v interface{}) {
			totalCount++

			connection := v.(*connectionm.Connection)
			if connection.WsConnectInfo.IsGuard() {
				guardCount++
			}
			if connection.WsConnectInfo.IsGuardMonth() {
				monthGuardCount++
			}
			if connection.WsConnectInfo.IsGuardYear() {
				yearGuardCount++
			}
			if connection.WsConnectInfo.IsRoomManager() {
				roomManagerCount++
			}
			if connection.WsConnectInfo.IsNobility() {
				nobilityCount++
			}
			if connection.WsConnectInfo.IsNobilityLevel1() {
				nobilityLevel1Count++
			}
			if connection.WsConnectInfo.IsNobilityLevel2() {
				nobilityLevel2Count++
			}
			if connection.WsConnectInfo.IsNobilityLevel3() {
				nobilityLevel3Count++
			}
			if connection.WsConnectInfo.IsNobilityLevel4() {
				nobilityLevel4Count++
			}
			if connection.WsConnectInfo.IsNobilityLevel5() {
				nobilityLevel5Count++
			}
			if connection.WsConnectInfo.IsNobilityLevel6() {
				nobilityLevel6Count++
			}
			if connection.WsConnectInfo.IsNobilityLevel7() {
				nobilityLevel7Count++
			}
		})

		id := serverID + "_" + liveID

		doc := &do.OnlineCountStatistic{
			ID:                  id,
			ServerID:            serverID,
			LiveID:              liveID,
			TotalCount:          totalCount,
			GuardCount:          guardCount,
			MonthGuardCount:     monthGuardCount,
			YearGuardCount:      yearGuardCount,
			NobilityCount:       nobilityCount,
			RoomManagerCount:    roomManagerCount,
			NobilityLevel1Count: nobilityLevel1Count,
			NobilityLevel2Count: nobilityLevel2Count,
			NobilityLevel3Count: nobilityLevel3Count,
			NobilityLevel4Count: nobilityLevel4Count,
			NobilityLevel5Count: nobilityLevel5Count,
			NobilityLevel6Count: nobilityLevel6Count,
			NobilityLevel7Count: nobilityLevel7Count,
		}
		docs = append(docs, doc)

	})

	// delete
	filter := bson.D{{"serverId", serverID}}
	_, err := dao.NewMgoOnlineCountStatisticDao().DeleteMany(filter)
	if err != nil {
		Log.Error(err)
		return
	}

	// insert
	if len(docs) > 0 {
		_, err = dao.NewMgoOnlineCountStatisticDao().InsertMany(docs)
		if err != nil {
			Log.Error(err)
			return
		}
	}
}

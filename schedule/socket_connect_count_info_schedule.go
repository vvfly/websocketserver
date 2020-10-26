package schedule

import (
	"strconv"

	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/model/dao"
	"github.com/luckyweiwei/websocketserver/model/do"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	"github.com/luckyweiwei/websocketserver/server/esort"
	"go.mongodb.org/mongo-driver/bson"
)

// socket 连接数统计刷新、直播间右上角在线用户排行榜列表刷新 SocketConnectCountInfo
func RefreshSocketConnectCountInfo() {
	Log.Debug("enter...")

	var (
		cm       = connectionm.GetConnectionManager()
		serverID = model.GetServerConfig().ServerID
		appID    = ""
		docs     = []interface{}{}
	)

	cm.LiveMap.IterCb(func(key string, v interface{}) {
		rank := make([]do.OnlineUser, 0)

		liveID := key
		roomConnection := v.(*connectionm.RoomConnection)
		if roomConnection == nil {
			Log.Warnf("cannot find roomID in map, liveID = %v", liveID)
			return
		}

		roomConnection.Connections.IterCb(func(key string, v interface{}) {
			connection := v.(*connectionm.Connection)

			// 同一直播间的所有用户在同一渠道下
			appID = connection.WsConnectInfo.Sid.AppID

			// 隐身进入不统计
			if !connection.WsConnectInfo.IsHide() {
				wsConnInfo := connection.WsConnectInfo
				rankInfo := do.OnlineUser{
					SdkVersion:         wsConnInfo.SdkVersion,
					AppID:              wsConnInfo.Sid.AppID,
					LiveAppID:          wsConnInfo.Authen.UsrInfo.AppIDForCurrentLive,
					SessionID:          wsConnInfo.SessionsInfoStr,
					IsReconnect:        wsConnInfo.Authen.UsrInfo.IsReconnect,
					UserID:             wsConnInfo.Sid.UserID,
					OpenID:             wsConnInfo.Authen.UsrInfo.OpenID,
					Avatar:             wsConnInfo.Authen.UsrInfo.Avatar,
					UserName:           wsConnInfo.Authen.UsrInfo.UserName,
					Sex:                wsConnInfo.Authen.UsrInfo.Sex,
					Role:               wsConnInfo.Authen.UsrInfo.Role,
					UserRole:           wsConnInfo.Authen.UsrInfo.UserRole,
					ExpGrade:           wsConnInfo.Authen.UsrInfo.ExpGrade,
					GuardType:          wsConnInfo.Authen.UsrInfo.GuardType,
					CarID:              wsConnInfo.Authen.UsrInfo.CarID,
					CarName:            wsConnInfo.Authen.UsrInfo.CarName,
					CarIcon:            wsConnInfo.Authen.UsrInfo.CarIcon,
					CarOnlineURL:       wsConnInfo.Authen.UsrInfo.CarOnlineURL,
					CarResURL:          wsConnInfo.Authen.UsrInfo.CarResURL,
					IsPlayCarAnim:      wsConnInfo.Authen.UsrInfo.IsPlayCarAnim,
					MarkUrlsJoinString: wsConnInfo.Authen.UsrInfo.MarkUrlsJoinString,
					NobilityType:       strconv.Itoa(wsConnInfo.Authen.UsrInfo.NobilityType),
					IsEnterHide:        wsConnInfo.Authen.UsrInfo.IsEnterHide,
					TokenType:          wsConnInfo.Authen.UsrInfo.TokenType,
				}

				rank = append(rank, rankInfo)
			}
		})

		count := len(rank)
		if count > 0 {
			rank = esort.SortRank(rank, proto.RankLength)

			id := serverID + "_" + appID + "_" + liveID

			doc := &do.SocketConnectCountInfo{
				ID:       id,
				ServerID: serverID,
				AppID:    appID,
				LiveID:   liveID,
				Count:    count,
				Rank:     rank,
			}
			docs = append(docs, doc)
		}
	})

	// delete
	filter := bson.D{{"serverId", serverID}}
	_, err := dao.NewMgoSocketConnectCountInfoDao().DeleteMany(filter)
	if err != nil {
		Log.Error(err)
		return
	}

	// insert
	if len(docs) > 0 {
		_, err = dao.NewMgoSocketConnectCountInfoDao().InsertMany(docs)
		if err != nil {
			Log.Error(err)
			return
		}
	}
}

package schedule

import (
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/model/dao"
	"github.com/luckyweiwei/websocketserver/model/do"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	"go.mongodb.org/mongo-driver/bson"
)

// 每个渠道在线用户数统计
func RefreshAppMemberCount() {
	Log.Debug("enter...")

	var (
		cm       = connectionm.GetConnectionManager()
		serverID = model.GetServerConfig().ServerID
		docs     = []interface{}{}
	)

	cm.AppMap.IterCb(func(key string, v interface{}) {
		appID := key
		appConnection := v.(*connectionm.AppConnection)
		if appConnection == nil {
			Log.Warnf("cannot find appID in map, appID = %v", appID)
			return
		}

		appMemCount := int64(appConnection.Connections.Count())

		if appMemCount > 0 {
			id := serverID + "_" + appID

			doc := &do.AppMemberCount{
				ID:       id,
				ServerID: serverID,
				AppID:    appID,
				Count:    appMemCount,
			}
			docs = append(docs, doc)
		}
	})

	// delete
	filter := bson.D{{"serverId", serverID}}
	_, err := dao.NewMgoAppMemberCountDao().DeleteMany(filter)
	if err != nil {
		Log.Error(err)
		return
	}

	// insert
	if len(docs) > 0 {
		_, err = dao.NewMgoAppMemberCountDao().InsertMany(docs)
		if err != nil {
			Log.Error(err)
			return
		}
	}
}

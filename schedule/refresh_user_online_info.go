package schedule

import (
	"time"

	"github.com/gomodule/redigo/redis"
	redisclient "github.com/luckyweiwei/base/cache/redis-client"
	"github.com/luckyweiwei/base/grmon"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto"
)

func genUserOnlineInfo(wsConnInfo *proto.WsConnectInfo) string {
	value := proto.UserOnlineInfo{
		SdkVersion:         wsConnInfo.SdkVersion,
		SdkType:            wsConnInfo.SdkType,
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
		NobilityType:       wsConnInfo.Authen.UsrInfo.NobilityType,
		IsEnterHide:        wsConnInfo.Authen.UsrInfo.IsEnterHide,
		TokenType:          wsConnInfo.Authen.UsrInfo.TokenType,
		Reconnect:          wsConnInfo.IsReconnect(),
		Hide:               wsConnInfo.IsHide(),
		PushSide:           wsConnInfo.IsPushSide(),
		Nobility:           wsConnInfo.IsNobility(),
		Guard:              wsConnInfo.IsGuard(),
		LiveID:             wsConnInfo.Sid.LiveID,
		EnterType:          wsConnInfo.Sid.EnterType,
		Login:              wsConnInfo.IsLogin(),
		PullSide:           wsConnInfo.IsPullSide(),
		RoomManager:        wsConnInfo.IsRoomManager(),
		Vip:                wsConnInfo.IsVip(),
		NotHide:            !wsConnInfo.IsHide(),
		IsRankHide:         wsConnInfo.Authen.UsrInfo.IsRankHide,
	}

	return utils.SerializeToJson(&value)
}

func MsgRefresh(wsConnInfo *proto.WsConnectInfo) {
	key := model.OnlineUserKey + wsConnInfo.Sid.UserID

	redisClient := redisclient.GetRedisClientManager().GetRedisClient(model.RedisDBNameDB1)
	conn := redisClient.Get()
	defer conn.Close()

	replyStr, err := redis.String(conn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		Log.Error(err)
		return
	}

	if replyStr == "" {
		valueJSONStr := genUserOnlineInfo(wsConnInfo)
		err := conn.Send("SET", key, valueJSONStr)
		if err != nil {
			Log.Error(err)
			return
		}
		err = conn.Send("EXPIRE", key, 100)
		if err != nil {
			Log.Error(err)
			return
		}
		conn.Flush()

		return
	}

	_, err = redis.Int64(conn.Do("EXPIRE", key, 100))
	if err != nil {
		Log.Error(err)
		return
	}
}

func ConnRefresh(wsConnInfo *proto.WsConnectInfo) {
	key := model.OnlineUserKey + wsConnInfo.Sid.UserID

	redisClient := redisclient.GetRedisClientManager().GetRedisClient(model.RedisDBNameDB1)
	conn := redisClient.Get()
	defer conn.Close()

	valueJSONStr := genUserOnlineInfo(wsConnInfo)
	err := conn.Send("SET", key, valueJSONStr)
	if err != nil {
		Log.Error(err)
		return
	}
	err = conn.Send("EXPIRE", key, 100)
	if err != nil {
		Log.Error(err)
		return
	}
	conn.Flush()
}

func DisConnRefresh(wsConnInfo *proto.WsConnectInfo) {
	key := model.OnlineUserKey + wsConnInfo.Sid.UserID

	redisClient := redisclient.GetRedisClientManager().GetRedisClient(model.RedisDBNameDB1)
	conn := redisClient.Get()
	defer conn.Close()

	_, err := redis.Int64(conn.Do("DEL", key))
	if err != nil {
		Log.Error(err)
		return
	}
}

func RefreshUserOnlineInfoConn() {
	grm := grmon.GetGRMon()

	for {
		select {
		case wsConnInfoConn := <-proto.UserOnlineInfoConnChan:
			connMsg := wsConnInfoConn
			grm.Go("ConnRefresh", func() { ConnRefresh(connMsg) })
		}
	}
}

func RefreshUserOnlineInfoDisConn() {
	// grm := grmon.GetGRMon()

	// for {
	// 	select {
	// 	case wsConnInfoConn := <-proto.UserOnlineInfoDisConnChan:
	// 		connMsg := wsConnInfoConn
	// 		grm.Go("DisConnRefresh", func() { DisConnRefresh(connMsg) })
	// 	}
	// }
}

func RefreshUserOnlineInfoMsg() {
	grm := grmon.GetGRMon()
	MsgMap := make(map[string]*proto.WsConnectInfo)

	lenUserOnlineInfoMsgChan := len(proto.UserOnlineInfoMsgChan)
	if lenUserOnlineInfoMsgChan > 0 {
		for i := 0; i < lenUserOnlineInfoMsgChan; i++ {
			wsConnInfoMsg := <-proto.UserOnlineInfoMsgChan
			MsgMap[wsConnInfoMsg.Sid.UserID] = wsConnInfoMsg
		}
	}

	for _, v := range MsgMap {
		msg := v
		grm.Go("RefreshUserOnlineInfoMsg", func() { MsgRefresh(msg) })
		time.Sleep(100 * time.Microsecond)
	}

	time.Sleep(30 * time.Second)
}

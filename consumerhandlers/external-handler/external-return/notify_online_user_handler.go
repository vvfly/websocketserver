package externalreturn

import (
	"errors"

	"github.com/Shopify/sarama"
	"github.com/gomodule/redigo/redis"
	redisclient "github.com/luckyweiwei/base/cache/redis-client"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
)

func NotifyOnlineUserHandler(msg *sarama.ConsumerMessage) error {
	Log.Debug("entering ...")

	userID := string(msg.Key)
	key := model.OnlineUserKey + userID

	redisClient := redisclient.GetRedisClientManager().GetRedisClient(model.RedisDBNameDB1)
	conn := redisClient.Get()
	defer conn.Close()

	replyStr, err := redis.String(conn.Do("GET", key))
	if err != nil {
		Log.Error(err)
		return err
	}

	value := &proto.UserOnlineInfo{}
	err = utils.DecodeFromJson(replyStr, value)
	if err != nil {
		Log.Error(err)
		return err
	}

	sessionID := value.SessionID
	cm := connectionm.GetConnectionManager()
	connection := cm.GetConnection(sessionID)
	if connection == nil {
		Log.Warnf("cannot find conn in map, sessionID = %v", sessionID)
		return errors.New("cannot find conn in map, sessionID =" + sessionID)
	}

	// 更新
	connection.WsConnectInfo.Authen.UsrInfo.IsReconnect = value.IsReconnect
	connection.WsConnectInfo.Authen.UsrInfo.Avatar = value.Avatar
	connection.WsConnectInfo.Authen.UsrInfo.Sex = value.Sex
	connection.WsConnectInfo.Authen.UsrInfo.Role = value.Role
	connection.WsConnectInfo.Authen.UsrInfo.UserRole = value.UserRole
	connection.WsConnectInfo.Authen.UsrInfo.ExpGrade = value.ExpGrade
	connection.WsConnectInfo.Authen.UsrInfo.GuardType = value.GuardType
	connection.WsConnectInfo.Authen.UsrInfo.CarID = value.CarID
	connection.WsConnectInfo.Authen.UsrInfo.CarName = value.CarName
	connection.WsConnectInfo.Authen.UsrInfo.CarIcon = value.CarIcon
	connection.WsConnectInfo.Authen.UsrInfo.CarOnlineURL = value.CarOnlineURL
	connection.WsConnectInfo.Authen.UsrInfo.CarResURL = value.CarResURL
	connection.WsConnectInfo.Authen.UsrInfo.IsPlayCarAnim = value.IsPlayCarAnim
	connection.WsConnectInfo.Authen.UsrInfo.MarkUrlsJoinString = value.MarkUrlsJoinString
	connection.WsConnectInfo.Authen.UsrInfo.NobilityType = value.NobilityType
	connection.WsConnectInfo.Authen.UsrInfo.IsEnterHide = value.IsEnterHide

	return nil
}

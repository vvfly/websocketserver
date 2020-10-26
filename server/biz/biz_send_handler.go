package biz

import (
	"errors"
	"strconv"

	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/proto/kafkaproto"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
)

func BizSendHandler(key string, req *proto.MsgReq) error {

	Log.Debug("entering ...")

	var (
		messageType = req.MessageType
	)

	// 用户信息
	cm := connectionm.GetConnectionManager()
	connection := cm.GetConnection(key)
	if connection == nil {
		Log.Warnf("cannot find user in map, sessionID = %v", key)
		return errors.New("cannot find user in map, sessionID =" + key)
	}
	userInfo := connection.WsConnectInfo

	// 填充消息
	businessData := req.BusinessData.(map[string]interface{})
	if messageType == proto.MT_CHAT {
		businessData["sex"] = userInfo.Authen.UsrInfo.Sex
		businessData["markUrlsJoinString"] = userInfo.Authen.UsrInfo.MarkUrlsJoinString
	}
	businessData["userName"] = userInfo.Authen.UsrInfo.UserName
	businessData["userId"] = userInfo.Authen.UsrInfo.UserID
	businessData["liveId"] = userInfo.Sid.LiveID
	businessData["role"] = userInfo.Authen.UsrInfo.Role
	businessData["expGrade"] = strconv.Itoa(userInfo.Authen.UsrInfo.ExpGrade)
	businessData["guardType"] = userInfo.Authen.UsrInfo.GuardType
	businessData["avatar"] = userInfo.Authen.UsrInfo.Avatar
	businessData["clientIp"] = userInfo.Authen.UsrInfo.ClientIP
	businessData["ks"] = strconv.FormatInt(utils.GetTimeOfS(), 10)

	req.BusinessData = businessData
	msgReqStr := utils.SerializeToJson(req)

	// 发送kafka
	produceData := kafkaproto.SetKfkProducerJobData(messageType, key, []byte(msgReqStr))
	kafkaproto.KafkaProducerJobChan <- produceData

	return nil
}

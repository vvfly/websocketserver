package biz

import (
	"errors"

	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/proto/kafkaproto"
	"github.com/luckyweiwei/websocketserver/proto/wproto"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
)

func LeaveHandler(key, value string) error {
	Log.Debug("entering ...")

	var (
		msgReq = &wproto.MsgLeaveReq{}
	)

	err := utils.DecodeFromJson(value, msgReq)
	if err != nil {
		Log.Error(err)
		return err
	}

	// 用户信息
	cm := connectionm.GetConnectionManager()
	connection := cm.GetConnection(key)
	if connection == nil {
		Log.Warnf("cannot find user in map, sessionID = %v", key)
		return errors.New("cannot find user in map, sessionID =" + key)
	}
	userInfo := connection.WsConnectInfo

	enterType := userInfo.Sid.EnterType
	if enterType != proto.EnterTypeAnchor {
		// 不是主播不处理
		return nil
	}

	// 填充消息 存入kafka
	msgWsReq := &wproto.MsgLeaveWsReq{
		R:           msgReq.R,
		S:           msgReq.S,
		MessageType: msgReq.MessageType,
		T:           msgReq.T,
	}
	msgWsReq.BusinessData = wproto.MsgLeaveWsReqData{
		Avatar:    userInfo.Authen.UsrInfo.Avatar,
		UserName:  userInfo.Authen.UsrInfo.UserName,
		LiveID:    userInfo.Sid.LiveID,
		Role:      userInfo.Authen.UsrInfo.Role,
		ExpGrade:  userInfo.Authen.UsrInfo.ExpGrade,
		GuardType: userInfo.Authen.UsrInfo.GuardType,
		ClientIP:  userInfo.Authen.UsrInfo.ClientIP,
		KS:        utils.GetTimeOfS(),
	}

	// 消息存入kafka
	topic := msgReq.MessageType
	wsValue := utils.SerializeToJson(msgWsReq)

	produceData := kafkaproto.SetKfkProducerJobData(topic, key, []byte(wsValue))
	kafkaproto.KafkaProducerJobChan <- produceData

	return nil
}

package biz

import (
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/helper"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto"
)

// 客户端消息签名验证
func VerifyClientSign(msgReq *proto.MsgReq) bool {

	if msgReq == nil {
		return false
	}

	r := msgReq.R
	s := msgReq.S
	t := msgReq.T
	businessData := utils.SerializeToJson(msgReq.BusinessData)

	serverConfig := model.GetServerConfig()
	yuanwen := businessData + t + r + serverConfig.SignatureSecretKey

	return helper.MD5(yuanwen) == s
}

// 客户端消息处理
func BizHandler(key, value string) error {
	Log.Debug("entering ...")

	msgReq := &proto.MsgReq{}
	err := utils.DecodeFromJson(value, msgReq)
	if err != nil {
		Log.Error(err)
		return err
	}

	// 客户端消息签名验证
	// if !VerifyClientSign(msgReq) {
	// 	Log.Error("客户端消息签名验证失败")
	// 	return nil
	// }

	messageType := msgReq.MessageType
	switch messageType {
	// 礼物服务
	case proto.MT_GIFT:
		err = GiftHandler(key, value)

	// 直播服务
	case proto.MT_LEAVE:
		err = LeaveHandler(key, value)

	default: // 通用客户端消息处理
		err = BizSendHandler(key, msgReq)
	}

	return err
}

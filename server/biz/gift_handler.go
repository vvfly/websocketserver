package biz

import (
	"errors"
	"strconv"

	httpclient "github.com/luckyweiwei/base/http-client"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/proto/wproto"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	"github.com/luckyweiwei/websocketserver/server/consul"
	remoteservice "github.com/luckyweiwei/websocketserver/server/remote-service"
)

func GiftHandler(key, value string) error {
	Log.Debug("entering ...")

	var (
		msgReq = &wproto.MsgGiftReq{}
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
	connInfo := connection.WsConnectInfo

	// 填充消息
	msgWsReq := &wproto.MsgGiftWsReq{
		R:           msgReq.R,
		S:           msgReq.S,
		MessageType: msgReq.MessageType,
		T:           msgReq.T,
	}
	msgWsReq.BusinessData = wproto.MsgGiftWsReqData{
		GiftCostType: msgReq.BusinessData.GiftCostType,
		MarkID:       msgReq.BusinessData.MarkID,
		GiftName:     msgReq.BusinessData.GiftName,
		Sex:          msgReq.BusinessData.Sex,
		BoxType:      msgReq.BusinessData.BoxType,
		AnchorID:     msgReq.BusinessData.AnchorID,
		IsStarGift:   msgReq.BusinessData.IsStarGift,
		AnchorName:   msgReq.BusinessData.AnchorName,
		UUID:         msgReq.BusinessData.UUID,
		EffectType:   msgReq.BusinessData.EffectType,
		LiveCount:    msgReq.BusinessData.LiveCount,
		CreateTime:   msgReq.BusinessData.CreateTime,
		Price:        msgReq.BusinessData.Price,
		AppID:        msgReq.BusinessData.AppID,
		GiftNum:      msgReq.BusinessData.GiftNum,
		FollowStatus: msgReq.BusinessData.FollowStatus,

		UserName:  connInfo.Authen.UsrInfo.UserName,
		UserID:    connInfo.Authen.UsrInfo.UserID,
		ExpGrade:  strconv.Itoa(connInfo.Authen.UsrInfo.ExpGrade),
		GuardType: connInfo.Authen.UsrInfo.GuardType,
		Role:      connInfo.Authen.UsrInfo.Role,
		LiveID:    connInfo.Sid.LiveID,
		ClientIP:  connInfo.Authen.UsrInfo.ClientIP,
		Avatar:    connInfo.Authen.UsrInfo.Avatar,
		Ks:        strconv.FormatInt(utils.GetTimeOfS(), 10),
	}

	msgWsReqStr := utils.SerializeToJson(msgWsReq)

	postBody := &wproto.PostBody{
		SessionID: key,
		Message:   msgWsReqStr,
	}
	postBodyData := utils.SerializeToJson(postBody)

	giftCostType := msgWsReq.BusinessData.GiftCostType
	addr := ""
	if giftCostType == "1" {
		addr = consul.GetBalanceAddr(proto.GiftServerName)
		addr += proto.GiftSendSuf
	} else if giftCostType == "2" {
		addr = consul.GetBalanceAddr(proto.ShopServerName)
		addr += proto.ScoreGiftSend
	} else {
		Log.Error("invalid type. giftCostType=%v", giftCostType)
		return errors.New("invalid type. giftCostType=" + giftCostType)
	}

	// post 发送
	httpclient.New().
		Post(addr).
		Timeout(remoteservice.TimeOut).
		Type("json").
		Send(postBodyData).
		End()

	return nil
}

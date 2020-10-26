package ownreturn

import (
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/proto/wproto"
	"github.com/luckyweiwei/websocketserver/server/broadcast"
)

func EntryReturnHandler(msg *sarama.ConsumerMessage) error {
	Log.Debug("entering ...")

	var (
		value = string(msg.Value)

		msgReturnValue = &wproto.MsgEntryWsReq{}

		msgResp         = &wproto.MsgEntryResp{}
		msgOrdinaryResp = &wproto.MsgOrdinaryEntryResp{}
		RespStr         = ""
	)

	err := utils.DecodeFromJson(value, msgReturnValue)
	if err != nil {
		Log.Error(err)
		return err
	}

	isOrdinary := msgReturnValue.IsdOrdinary
	if isOrdinary == 1 { // 普通用户
		msgOrdinaryResp.MessageType = proto.MT_ENTRY
		msgOrdinaryResp.BusinessData.Code = proto.SUCCESS
		msgOrdinaryResp.BusinessData.ResultData = wproto.MsgOrdinaryEntryRespData{
			UserName: msgReturnValue.UserName,
			ExpGrade: msgReturnValue.ExpGrade,
		}

		RespStr = utils.SerializeToJson(msgOrdinaryResp)

	} else { // 非普通用户
		msgResp.MessageType = proto.MT_ENTRY
		msgResp.BusinessData.Code = proto.SUCCESS

		markUrls := strings.Split(msgReturnValue.MarkUrlsJoinString, "|")

		msgResp.BusinessData.ResultData = wproto.MsgEntryRespData{
			UserName:                     msgReturnValue.UserName,
			UserID:                       msgReturnValue.UserID,
			Role:                         msgReturnValue.Role,
			UserRole:                     msgReturnValue.UserRole,
			Sex:                          msgReturnValue.Sex,
			Avatar:                       msgReturnValue.Avatar,
			ExpGrade:                     msgReturnValue.ExpGrade,
			GuardType:                    msgReturnValue.GuardType,
			CarID:                        msgReturnValue.CarID,
			CarName:                      msgReturnValue.CarName,
			CarIcon:                      msgReturnValue.CarIcon,
			CarOnlineURL:                 msgReturnValue.CarOnlineURL,
			CarResURL:                    msgReturnValue.CarResURL,
			IsPlayCarAnim:                msgReturnValue.IsPlayCarAnim,
			NobilityType:                 msgReturnValue.NobilityType,
			IsEnterHide:                  msgReturnValue.IsEnterHide,
			IsPlayNobilityEnterAnimation: strconv.Itoa(msgReturnValue.IsPlayNobilityEnterAnimation),
			IsWeekStar:                   msgReturnValue.IsWeekStar,

			//加属性，用于渠道方app,用户详情跳转
			AppID:    msgReturnValue.AppID,
			OpenID:   msgReturnValue.OpenID,
			MarkUrls: markUrls,
		}

		RespStr = utils.SerializeToJson(msgResp)
	}

	// 广播
	broadcast.BroadcastEntry(msgReturnValue.LiveID, RespStr, string(msg.Key))

	return nil
}

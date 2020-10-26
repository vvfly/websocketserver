package externalreturn

import (
	"errors"

	"github.com/Shopify/sarama"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/broadcast"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
)

func ExternalReturnHandler(msg *sarama.ConsumerMessage) error {
	Log.Debug("entering ...")

	var (
		key   = string(msg.Key)
		value = string(msg.Value)

		msgReturnKey = &proto.MsgReturnKey{}
	)

	err := utils.DecodeFromJson(key, msgReturnKey)
	if err != nil {
		Log.Error(err)
		return err
	}

	messageType := msgReturnKey.MessageType
	if messageType == proto.MT_TOKEN_INVALID_NOTIFY { // 关闭通道
		cm := connectionm.GetConnectionManager()
		connection := cm.GetConnection(key)
		if connection == nil {
			Log.Warnf("cannot find conn in map, sessionID = %v", key)
			return errors.New("cannot find conn in map")
		}
		cm.ClearConnection(connection)
		return nil
	}

	scopeType := msgReturnKey.ScopeType
	if scopeType == proto.BY_BULK {
		return broadcastByBulk(msgReturnKey, value)
	}

	return broadcastBySingle(msgReturnKey, value)
}

func broadcastByBulk(msgReturnKey *proto.MsgReturnKey, value string) error {
	Log.Debug("enter...")

	var (
		msgReturnValueBulk = &proto.MsgReturnValueBulk{}
	)

	err := utils.DecodeFromJson(value, &msgReturnValueBulk.MsgReturnValueBulkData)
	if err != nil {
		Log.Error(err)
		return err
	}

	for _, v := range msgReturnValueBulk.MsgReturnValueBulkData {
		msgResp := &proto.MsgResp{}
		text := v.Text
		msgResp.MessageType = msgReturnKey.MessageType
		msgResp.BusinessData.Code = msgReturnKey.Code
		msgResp.BusinessData.Message = msgReturnKey.CodeMessage
		msgResp.BusinessData.ResultData = text

		msgRespStr := utils.SerializeToJson(msgResp)

		scopeType := v.ScopeType
		scopeID := v.ScopeID
		switch scopeType {
		case proto.BY_LIVE_ID:
			broadcast.BroadcastByLiveID(scopeID, msgRespStr)
		case proto.BY_APP_ID:
			broadcast.BroadcastByAppID(scopeID, msgRespStr)
		case proto.BY_SESSION_ID:
			broadcast.BroadcastBySessionID(scopeID, msgRespStr)
		case proto.BY_ALL:
			broadcast.BroadcastByAll(msgRespStr)
		default:
			Log.Warningf("Unknown scopeType. scopeType=%v", scopeType)
			continue
		}
	}

	return nil
}

func broadcastBySingle(msgReturnKey *proto.MsgReturnKey, value string) error {
	Log.Debug("entering ...")

	var (
		msgReturnValue = make(map[string]interface{})
		msgResp        = &proto.MsgResp{}
	)

	err := utils.DecodeFromJson(value, &msgReturnValue)
	if err != nil {
		Log.Error(err)
		return err
	}

	msgResp.MessageType = msgReturnKey.MessageType
	msgResp.BusinessData.Code = msgReturnKey.Code
	msgResp.BusinessData.Message = msgReturnKey.CodeMessage
	msgResp.BusinessData.ResultData = msgReturnValue

	msgRespStr := utils.SerializeToJson(msgResp)

	scopeType := msgReturnKey.ScopeType
	switch scopeType {
	case proto.BY_LIVE_ID, proto.BY_LIVE_IDS:
		broadcast.BroadcastByLiveIDs(msgReturnKey.ScopeIDList, msgRespStr)
	case proto.BY_APP_ID, proto.BY_APP_IDS:
		broadcast.BroadcastByAppIDs(msgReturnKey.ScopeIDList, msgRespStr)
	case proto.BY_SESSION_ID:
		broadcast.BroadcastBySessionIDs(msgReturnKey.ScopeIDList, msgRespStr)
	case proto.BY_ALL:
		broadcast.BroadcastByAll(msgRespStr)
	default:
		Log.Warningf("Unknown scopeType. scopeType=%v", scopeType)
		return errors.New("Unknown scopeType")
	}

	return nil
}

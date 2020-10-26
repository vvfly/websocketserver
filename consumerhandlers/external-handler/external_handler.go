package externalhandler

import (
	"errors"

	"github.com/Shopify/sarama"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	externalreturn "github.com/luckyweiwei/websocketserver/consumerhandlers/external-handler/external-return"
	ownreturn "github.com/luckyweiwei/websocketserver/consumerhandlers/external-handler/own-return"
	"github.com/luckyweiwei/websocketserver/proto/kafkaproto"
)

type ExternalHandler struct {
}

var externalHandler *ExternalHandler

func ExternalHandlerInit() *ExternalHandler {
	externalHandler = &ExternalHandler{}

	return externalHandler
}

func GetExternalHandler() *ExternalHandler {
	utils.ASSERT(externalHandler != nil)
	return externalHandler
}

func (external *ExternalHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
	if msg == nil {
		Log.Error("kafkaconsumer msg nil")
		return errors.New("kafkaconsumer msg nil")
	}

	var err error

	topic := msg.Topic
	switch topic {
	case kafkaproto.TOPIC_ENTRY:
		err = ownreturn.EntryReturnHandler(msg)
	case kafkaproto.TOPIC_NOTIFY_ONLINE_USER:
		err = externalreturn.NotifyOnlineUserHandler(msg)

	default: // 通用 return消息
		err = externalreturn.ExternalReturnHandler(msg)
	}

	return err
}

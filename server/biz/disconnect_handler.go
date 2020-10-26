package biz

import (
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/proto/kafkaproto"
)

func DisconnectHandler(topic, key, value string) error {
	Log.Debug("entering ...")

	produceData := kafkaproto.SetKfkProducerJobData(topic, key, []byte(value))
	kafkaproto.KafkaProducerJobChan <- produceData

	return nil
}

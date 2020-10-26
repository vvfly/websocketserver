package kafkatest

import (
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/kafka/consumer"
	"github.com/luckyweiwei/websocketserver/kafka/producer"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/proto/kafkaproto"
)

func TestProducer(t *testing.T) {
	// init
	serverConfig := model.GetServerConfig()
	_, err := toml.DecodeFile("../../config/websocket_server.toml", serverConfig)
	if err != nil {
		Log.Error(err)
		return
	}

	producer.InitKafkaProducer()

	msgReq := &proto.MsgReq{}

	produceEntryData := kafkaproto.SetKfkProducerJobData("chat_return", "xxx", []byte(utils.SerializeToJson(msgReq)))
	kafkaproto.KafkaProducerJobChan <- produceEntryData

	time.Sleep(5 * time.Second)
}

func TestConsumer(t *testing.T) {
	// init
	serverConfig := model.GetServerConfig()
	_, err := toml.DecodeFile("../../config/websocket_server.toml", serverConfig)
	if err != nil {
		Log.Error(err)
		return
	}

	consumer.InitKafkaConsumer()

	for {
		time.Sleep(5 * time.Second)
	}
}

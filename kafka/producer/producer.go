package producer

import (
	"github.com/luckyweiwei/base/grmon"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto/kafkaproto"

	"github.com/Shopify/sarama"
	. "github.com/luckyweiwei/base/logger"
)

type SimAsyncProducer struct {
	asyncProducer sarama.AsyncProducer
}

func (sap *SimAsyncProducer) Start() {
	grm := grmon.GetGRMon()
	grm.Go("SimAsyncProducer Work", sap.work)
}

func (sap *SimAsyncProducer) Stop() {
	Log.Info("stop SimAsyncProducer.")
	sap.asyncProducer.AsyncClose()
}

func (sap *SimAsyncProducer) work() {
	defer sap.Stop()

	grm := grmon.GetGRMon()
	for {
		select {
		case success := <-sap.asyncProducer.Successes():
			key, _ := success.Key.Encode()
			value, _ := success.Value.Encode()
			Log.Infof("Producer produce message success...\ntopic=%v\nkey=%v\nvalue=%v", success.Topic, string(key), string(value))
		case err := <-sap.asyncProducer.Errors():
			Log.Errorf("asyncProducer err:%v\n", err)
		case msg := <-kafkaproto.KafkaProducerJobChan:
			produceMsg := msg
			grm.Go("ProduceMessage", func() { sap.produce(produceMsg) })
		}
	}
}

func (sap *SimAsyncProducer) produce(record *kafkaproto.KfkProducerJobData) {

	serverConfig := model.GetServerConfig()

	topic := record.MsgType
	if topic == kafkaproto.TOPIC_BIZ {
		topic = serverConfig.ServerID + "_" + topic
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(record.Key),
		Value: sarama.ByteEncoder(record.Record),
	}

	Log.Debugf("Message begin to write into kafka. message = %v", utils.FormatStruct(message))

	sap.asyncProducer.Input() <- message
}

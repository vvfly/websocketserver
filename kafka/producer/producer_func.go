package producer

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/model"
)

func NewProducerConfiguration() *sarama.Config {
	// serverConfig := model.GetServerConfig()
	config := sarama.NewConfig()

	//socket.timeout.ms.  default 30 seconds
	config.Net.DialTimeout = 5000 * time.Millisecond
	config.Net.WriteTimeout = 5000 * time.Millisecond
	config.Net.ReadTimeout = 5000 * time.Millisecond

	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	// config.Producer.Flush.Messages = serverConfig.ProducerConf.BatchNumer
	config.Producer.Flush.Frequency = time.Millisecond * 100
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.NoResponse
	// 使用snappy压缩
	config.Producer.Compression = sarama.CompressionSnappy
	return config
}

func NewSimAsyncProducer() (producer *SimAsyncProducer, err error) {
	serverConfig := model.GetServerConfig()

	producer = &SimAsyncProducer{}
	producer.asyncProducer, err = sarama.NewAsyncProducer(serverConfig.ProducerConf.BrokerList, NewProducerConfiguration())
	if err != nil {
		Log.Errorf("Failed to NewSimAsyncProducer. err = %v", err)
		return nil, err
	}
	return producer, nil
}

func InitKafkaProducer() {
	producer, err := NewSimAsyncProducer()
	if err != nil {
		Log.Errorf("NewSimAsyncProducer failed. err = %v", err)
		os.Exit(1)
	}
	producer.Start()
}

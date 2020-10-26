package consumer

import (
	"os"
	"regexp"
	"strings"

	"github.com/Shopify/sarama"
	. "github.com/luckyweiwei/base/logger"
	externalhandler "github.com/luckyweiwei/websocketserver/consumerhandlers/external-handler"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto/kafkaproto"
)

type ConsumerOptions struct {
	BrokerAddresses []string
	Topics          []string
	GroupID         string
	Version         string

	Cfg *sarama.Config
}

func NewConsumerOptions() *ConsumerOptions {
	return &ConsumerOptions{}
}

func arrayToString(arr []string) string {
	var result string
	for _, v := range arr {
		result += v
	}
	return result
}

func SetConsumerOption(cfg *model.KafkaConsumerConfig, opts *ConsumerOptions) error {
	serverConfig := model.GetServerConfig()

	// topics
	topics := cfg.Topics
	for _, topic := range topics {
		if topic == kafkaproto.TOPIC_BIZ {
			topic = serverConfig.ServerID + "_" + topic
		}

		opts.Topics = append(opts.Topics, topic)
	}

	opts.GroupID = serverConfig.ServerID + "_" + cfg.GroupID
	opts.BrokerAddresses = cfg.BrokerAddresses
	opts.Version = cfg.Version

	opts.Cfg = sarama.NewConfig()

	version, err := sarama.ParseKafkaVersion(opts.Version)
	if err != nil {
		Log.Error(err)
		return err
	}
	opts.Cfg.Version = version
	opts.Cfg.Consumer.Return.Errors = true
	opts.Cfg.Consumer.Offsets.Initial = sarama.OffsetNewest

	return opts.Cfg.Validate()
}

func SetReturnTopicOption(tag string, opts *ConsumerOptions) error {
	if tag == model.TAG_EXTERNAL_FOO {
		opts.Topics = append(opts.Topics, kafkaproto.FooTopics...)

		ca, err := sarama.NewClusterAdmin(opts.BrokerAddresses, opts.Cfg)
		if err != nil {
			Log.Error(err)
			return err
		}
		defer ca.Close()

		topicDetail, err := ca.ListTopics()
		if err != nil {
			Log.Error(err)
			return err
		}

		reg, err := regexp.Compile("_return$")
		if err != nil {
			Log.Error(err)
			return err
		}

		fooTopicsStr := arrayToString(kafkaproto.FooTopics)
		for topic := range topicDetail {
			if reg.MatchString(topic) {
				if !strings.Contains(fooTopicsStr, topic) {
					opts.Topics = append(opts.Topics, topic)
				}
			}
		}
	} else if tag == model.TAG_EXTERNAL_BAR {
		opts.Topics = append(opts.Topics, kafkaproto.BarTopics...)

		ca, err := sarama.NewClusterAdmin(opts.BrokerAddresses, opts.Cfg)
		if err != nil {
			Log.Error(err)
			return err
		}
		defer ca.Close()

		topicDetail, err := ca.ListTopics()
		if err != nil {
			Log.Error(err)
			return err
		}

		reg, err := regexp.Compile("_return$")
		if err != nil {
			Log.Error(err)
			return err
		}

		barTopicsStr := arrayToString(kafkaproto.BarTopics)
		for topic := range topicDetail {
			if reg.MatchString(topic) {
				if !strings.Contains(barTopicsStr, topic) {
					opts.Topics = append(opts.Topics, topic)
				}
			}
		}
	}

	Log.Debugf("Tag=%v, topics=%v", tag, opts.Topics)

	return nil
}

func InitKafkaConsumer() {
	//加载consumer配置
	serverConfig := model.GetServerConfig()

	for tag, cfg := range serverConfig.ConsumerConf {
		opts := NewConsumerOptions()
		err := SetConsumerOption(&cfg, opts)
		if err != nil {
			Log.Error(err)
			os.Exit(1)
		}

		err = SetReturnTopicOption(tag, opts)
		if err != nil {
			Log.Error(err)
			os.Exit(1)
		}

		consumerGroup := NewSimConsumerGroup(opts)

		var handler IHandler
		switch tag {
		case model.TAG_EXTERNAL_FOO, model.TAG_EXTERNAL_BAR:
			handler = externalhandler.GetExternalHandler()

		default:
			Log.Warningf("Unknown tag: %s\n", tag)
			os.Exit(1)
		}

		consumerGroup.SetHandler(handler)

		//启动kafka消费
		err = consumerGroup.Start()
		if err != nil {
			Log.Errorf("Consumer Start failed;err:%s\n", err)
			os.Exit(1)
		}
	}
}

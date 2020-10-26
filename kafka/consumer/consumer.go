package consumer

import (
	"context"
	"errors"
	"time"

	"github.com/Shopify/sarama"
	"github.com/luckyweiwei/base/grmon"
	. "github.com/luckyweiwei/base/logger"
)

type IHandler interface {
	HandleMessage(message *sarama.ConsumerMessage) error
	// TimeUsed() time.Duration
	// Exit()
}

type SimConsumerGroup struct {
	opts      *ConsumerOptions
	consumer  sarama.ConsumerGroup
	ctx       context.Context
	cancel    context.CancelFunc
	msgHandle IHandler
}

func NewSimConsumerGroup(opts *ConsumerOptions) *SimConsumerGroup {
	c := &SimConsumerGroup{
		opts: opts,
	}
	return c
}

func (scg *SimConsumerGroup) SetHandler(handler IHandler) {
	scg.msgHandle = handler
}

func (scg *SimConsumerGroup) Start() error {
	if scg.msgHandle == nil {
		Log.Error("must set msg Handler")
		return errors.New("must set msg Handler")
	}

	ctx, cancel := context.WithCancel(context.Background())
	consumer, err := sarama.NewConsumerGroup(scg.opts.BrokerAddresses, scg.opts.GroupID, scg.opts.Cfg)
	if err != nil {
		Log.Errorf("NewConsumerGroup err: %s\n", err)
		cancel()
		return err
	}
	scg.consumer = consumer
	scg.ctx = ctx
	scg.cancel = cancel

	grm := grmon.GetGRMon()
	grm.Go("dealErrors", scg.DealErrors)
	grm.Go("handleLoop", scg.HandleLoop)

	return nil
}

func (scg *SimConsumerGroup) DealErrors() {
	Log.Info("Consumer errorloop begin.")

	defer scg.Stop()

	for {
		select {
		case err := <-scg.consumer.Errors():
			Log.Warningf("consume error. err = %v", err)
			time.Sleep(5 * time.Second)
		}
	}
}

func (scg *SimConsumerGroup) HandleLoop() {
	Log.Info("Consumer handleloop begin.")

	defer scg.Stop()

	for {
		err := scg.consumer.Consume(scg.ctx, scg.opts.Topics, scg)
		if err != nil {
			Log.Error(err)
			// 5秒后重试
			time.Sleep(time.Second)
		}
	}
}

func (scg *SimConsumerGroup) Stop() {
	Log.Info("stop SimConsumerGroup.")
	scg.cancel()
	scg.consumer.Close()
}

func (scg *SimConsumerGroup) Setup(s sarama.ConsumerGroupSession) error {
	Log.Debug("consumer begin...")
	return nil
}

func (scg *SimConsumerGroup) Cleanup(s sarama.ConsumerGroupSession) error {
	Log.Debug("consumer end...")
	return nil
}

func (scg *SimConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	grm := grmon.GetGRMon()

	for msg := range claim.Messages() {
		Log.Infof("begin handle msg...\ntopic=%v\nkey=%v\nvalue=%v\npartition=%v\noffset=%v\n",
			msg.Topic, string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)

		// 异步处理消息
		handleMsg := msg // 避免闭包
		grm.Go("HandleMessage", func() { scg.msgHandle.HandleMessage(handleMsg) })
		session.MarkMessage(handleMsg, "")
		time.Sleep(time.Nanosecond)
	}

	return nil
}

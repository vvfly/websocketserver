package kafkaproto

import "runtime"

var (
	KafkaProducerJobChan = make(chan *KfkProducerJobData, 1024*64*runtime.NumCPU())
)

type KfkProducerJobData struct {
	MsgType string
	Key     string
	Record  []byte
}

func SetKfkProducerJobData(topic, key string, record []byte) *KfkProducerJobData {
	data := &KfkProducerJobData{
		MsgType: topic,
		Key:     key,
		Record:  record,
	}

	return data
}

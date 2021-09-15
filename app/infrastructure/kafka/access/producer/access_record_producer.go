package producer

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
)

type AccessRecordKafkaProducer  struct {
	producer sarama.SyncProducer
}

const (
	BrokersUrl         = "localhost:9092"
	AccessRecordTopics = "access-record"
)

func NewAccessRecordKafkaProducer() *AccessRecordKafkaProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
	conn, err := sarama.NewSyncProducer(strings.Split(BrokersUrl,","), config)
	if err != nil {
		return nil
	}
	
	return &AccessRecordKafkaProducer{producer: conn}
}

func (c AccessRecordKafkaProducer) Close()  {
	err := c.producer.Close()
	if err != nil {
		return
	}
}

func (c *AccessRecordKafkaProducer)PushMessage(message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: AccessRecordTopics,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := c.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", AccessRecordTopics, partition, offset)
	return nil
}

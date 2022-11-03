package events

import (
	"fmt"
	"github/Services/apuc/userservice/config"
	"github/Services/apuc/userservice/pkg/logger"
	broker "github/Services/apuc/userservice/pkg/messagebroker"
	"time"
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	kafkaWriter *kafka.Writer
	log         logger.Logger
}

func NewKafkaProducerBroker(conf config.Config, log logger.Logger, topic string) broker.Publisher {
	connString := fmt.Sprintf("%s:%d", conf.KafkaHost, conf.KafkaPort)

	return &KafkaPublisher{
		kafkaWriter: &kafka.Writer{
			Addr:         kafka.TCP(connString),
			Topic:        topic,
			BatchTimeout: 10 * time.Millisecond,
		},
		log: log,
	}
}

func (p *KafkaPublisher) Start() error {
	return nil
}
func (p *KafkaPublisher) Stop() error {
	err := p.kafkaWriter.Close()
	if err != nil {
		return err
	}
	return nil
}
func (p *KafkaPublisher) Publish(key, body []byte, logBody string) error {
	message := kafka.Message{
		Key:   key,
		Value: body,
	}

	if err := p.kafkaWriter.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	
	p.log.Info("Message published(key/body): " + string(key) + "/" + logBody)
	return nil
}
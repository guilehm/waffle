package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log/slog"
	"tmdb/pkg/logging"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(brokers string, maxMessages int) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            brokers,
		"queue.buffering.max.messages": maxMessages,
		"queue.buffering.max.kbytes":   1024000,
		"linger.ms":                    0,
		"acks":                         "all",
	})
	if err != nil {
		return nil, err
	}
	producer := &Producer{producer: p}
	go producer.handleDeliveryReports()
	return producer, nil
}

func (p *Producer) handleDeliveryReports() {
	logger := logging.Logger
	for e := range p.producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				logger.Error(
					"failed to deliver message",
					slog.String("topic_partition", fmt.Sprintf("%v", ev.TopicPartition)),
				)
				continue
			}

			logger.Info(
				"successfully produced message",
				slog.String("topic", *ev.TopicPartition.Topic),
				slog.Int("partition", int(ev.TopicPartition.Partition)),
				slog.String("offset", ev.TopicPartition.Offset.String()),
			)
		}
	}
}

func (p *Producer) Produce(topic string, message []byte) error {
	logger := logging.Logger

	err := p.producer.Produce(&kafka.Message{
		// TODO: set the key
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)
	if err != nil {
		logger.Error("failed to send message to Kafka", slog.String("error", err.Error()))
		return err
	}
	return nil
}

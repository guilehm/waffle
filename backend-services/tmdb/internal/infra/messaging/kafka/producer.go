package kafka

import (
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
	})
	if err != nil {
		return nil, err
	}
	return &Producer{producer: p}, nil
}

func (p *Producer) Publish(topic string, message []byte) error {
	logger := logging.Logger
	deliveryChan := make(chan kafka.Event, 1)

	err := p.producer.Produce(&kafka.Message{
		// TODO: set the key
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)
	if err != nil {
		logger.Error("failed to send message to Kafka", slog.String("error", err.Error()))
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		logger.Error("delivery failed", slog.String("error", m.TopicPartition.Error.Error()))
		return m.TopicPartition.Error
	}

	logger.Info(
		"successfully delivered message",
		slog.String("topic", topic),
		slog.Int("partition", int(m.TopicPartition.Partition)),
		slog.String("offset", m.TopicPartition.Offset.String()),
	)
	return nil
}

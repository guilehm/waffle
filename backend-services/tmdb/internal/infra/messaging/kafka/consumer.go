package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log/slog"
	"tmdb/pkg/logging"
)

type Consumer struct {
	consumer *kafka.Consumer
}

func NewConsumer(brokers, groupID string) (*Consumer, error) {
	logger := logging.Logger
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		logger.Error("failed to create consumer", slog.String("error", err.Error()))
		return nil, err
	}
	return &Consumer{consumer: c}, nil
}

func (c *Consumer) Consume(topics []string, handleMessage func(*kafka.Message)) error {
	logger := logging.Logger
	err := c.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		logger.Error("failed to subscribe to topics", slog.String("error", err.Error()))
		return err
	}

	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logger.Error("error while reading message", slog.String("error", err.Error()))
			continue
		}
		handleMessage(msg)
	}
}

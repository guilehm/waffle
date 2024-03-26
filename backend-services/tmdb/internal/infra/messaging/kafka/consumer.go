package kafka

import (
	"fmt"
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
		"bootstrap.servers":               brokers,
		"group.id":                        groupID,
		"auto.offset.reset":               "earliest",
		"go.application.rebalance.enable": true,
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
		logger.Error(
			"failed to subscribe to topics",
			slog.String("error", err.Error()),
		)
		return err
	}

	run := true
	for run == true {
		ev := c.consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			logger.Info(
				"message consumed",
				slog.String("topic_partition", fmt.Sprintf("%v", e.TopicPartition)),
			)
		case kafka.Error:
			logger.Error(
				"error while consuming message",
				slog.String("error", e.Error()),
			)
			run = false
		default:
			logger.Debug(
				"ignored",
				slog.String("event", fmt.Sprintf("%v", e)),
			)
		}
	}
	return nil
}

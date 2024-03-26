package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log/slog"
	"tmdb/pkg/logging"
)

type Consumer struct {
	consumer       *kafka.Consumer
	minCommitCount int
}

func NewConsumer(brokers, groupID string, minCommitCount int) (*Consumer, error) {
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
	return &Consumer{consumer: c, minCommitCount: minCommitCount}, nil
}

func (c *Consumer) Consume(topics []string) error {
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
	msgCount := 0
	for run == true {
		ev := c.consumer.Poll(100)
		switch e := ev.(type) {
		case kafka.AssignedPartitions:
			logger.Info(
				"rebalance: assigned partitions",
				slog.String("event", fmt.Sprintf("%v", e)),
			)
			_ = c.consumer.Assign(e.Partitions)
		case kafka.RevokedPartitions:
			logger.Info(
				"rebalance: revoked partitions",
				slog.String("event", fmt.Sprintf("%v", e)),
			)
			_ = c.consumer.Unassign()
		case *kafka.Message:
			msgCount += 1
			if msgCount%c.minCommitCount == 0 {
				_, _ = c.consumer.Commit()
			}
			logger.Info(
				"message consumed",
				slog.String("topic_partition", fmt.Sprintf("%v", e.TopicPartition)),
			)
		case kafka.PartitionEOF:
			logger.Info(
				"reached end of partition",
				slog.String("event", fmt.Sprintf("%v", e)),
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

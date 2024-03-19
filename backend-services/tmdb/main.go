package main

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"tmdb/internal/app/usecase"
	"tmdb/internal/config"
	"tmdb/internal/infra/api"
	"tmdb/internal/infra/messaging/kafka"
	"tmdb/internal/infra/server"
	"tmdb/internal/infra/service"
	"tmdb/pkg/events"
)

func main() {
	fmt.Println("hello world")

	cfg := config.LoadConfig()

	// create kafka producer
	kafkaProducer, err := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaProducerMaxMessages)
	if err != nil {
		log.Fatalf("could not create kafka producer: %v", err)
	}

	// create kafka consumer
	consumer, err := kafka.NewConsumer(cfg.KafkaBrokers, events.MovieSearchTopic)
	if err != nil {
		log.Fatalf("could not create kafka consumer: %v", err)
	}

	// start consuming messages
	go func() {
		err = consumer.Consume([]string{events.MovieSearchTopic}, func(m *ckafka.Message) {
			fmt.Println("message consumed", m.TopicPartition)
		})
		if err != nil {
			log.Fatalf("could not start consumer: %v", err)
		}
	}()

	movieService := service.NewTMDBAPIClient(cfg.APIKey, cfg.APITimeout)
	movieUseCase := usecase.NewMovieUseCase(movieService, kafkaProducer)

	router := api.SetupRouter(movieUseCase)
	s, listener, err := server.CreateServer(router, cfg.Port)
	if err != nil {
		log.Fatalf("could not create server: %v", err)
	}

	err = s.Serve(*listener)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

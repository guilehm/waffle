package main

import (
	"fmt"
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
	consumer, err := kafka.NewConsumer(cfg.KafkaBrokers, cfg.AppName, cfg.KafkaMinCommitCount)
	if err != nil {
		log.Fatalf("could not create kafka consumer: %v", err)
	}

	// start consumer
	go func() {
		err = consumer.Consume([]string{events.MovieSearchTopic})
		if err != nil {
			log.Fatalf("could not consume messages: %v", err)
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

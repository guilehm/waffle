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
)

func main() {
	fmt.Println("hello world")

	cfg := config.LoadConfig()

	kafkaProducer, err := kafka.NewProducer(cfg.KafkaBrokers)
	if err != nil {
		log.Fatalf("could not create kafka producer: %v", err)
	}

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

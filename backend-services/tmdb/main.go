package main

import (
	"fmt"
	"log"
	"tmdb/internal/app/usecase"
	"tmdb/internal/config"
	"tmdb/internal/infra/api"
	"tmdb/internal/infra/server"
	"tmdb/internal/infra/service"
)

func main() {
	fmt.Println("hello world")

	cfg := config.LoadConfig()
	movieService := service.NewTMDBAPIClient(cfg.APIKey, cfg.APITimeout)
	movieUseCase := usecase.NewMovieUseCase(movieService)

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

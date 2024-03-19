package usecase

import (
	"encoding/json"
	"tmdb/internal/app/domain"
	"tmdb/internal/app/ports"
	"tmdb/pkg/events"
)

type MovieUseCase struct {
	tmdbService ports.TMDBService
	producer    ports.Messaging
}

func NewMovieUseCase(tmdbService ports.TMDBService, producer ports.Messaging) *MovieUseCase {
	return &MovieUseCase{
		tmdbService: tmdbService,
		producer:    producer,
	}
}

func (uc *MovieUseCase) MovieDetails(id string) (*domain.Movie, error) {
	return uc.tmdbService.MovieByID(id)
}

func (uc *MovieUseCase) FindMovies(query string, page int) (*domain.MovieSearchResponse, error) {
	topic := events.MovieSearchTopic
	response, err := uc.tmdbService.SearchMovies(query, page)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	err = uc.producer.Publish(topic, data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

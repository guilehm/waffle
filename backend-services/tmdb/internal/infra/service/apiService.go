package service

import (
	"net/http"
	"tmdb/internal/app/domain"
	"tmdb/internal/app/ports"
)

type TMDBAPIClient struct {
	baseURL string
	client  *http.Client
	apiKey  string
}

func NewMovieAPIClient(baseURL, apiKey string) ports.TMDBService {
	return &TMDBAPIClient{
		baseURL: baseURL,
		client:  &http.Client{},
		apiKey:  apiKey,
	}
}

func (c *TMDBAPIClient) MovieByID(id string) (*domain.Movie, error) {
	// make request, decode response to domain.Movie
	return nil, nil
}

func (c *TMDBAPIClient) SearchMovies(query string) ([]*domain.Movie, error) {
	return nil, nil
}

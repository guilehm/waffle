package service

import (
	"net/http"
	"tmdb/internal/app/domain"
	"tmdb/internal/app/ports"
)

type MovieAPIClient struct {
	baseURL string
	client  *http.Client
	apiKey  string
}

func NewMovieAPIClient(baseURL, apiKey string) ports.TMDBService {
	return &MovieAPIClient{
		baseURL: baseURL,
		client:  &http.Client{},
		apiKey:  apiKey,
	}
}

func (c *MovieAPIClient) MovieByID(id string) (*domain.Movie, error) {
	return nil, nil
}

func (c *MovieAPIClient) SearchMovies(query string) ([]*domain.Movie, error) {
	return nil, nil
}

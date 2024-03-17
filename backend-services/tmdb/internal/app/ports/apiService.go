package ports

import "tmdb/internal/app/domain"

type TMDBService interface {
	MovieByID(id string) (*domain.Movie, error)
	SearchMovies(query string) ([]*domain.Movie, error)
}

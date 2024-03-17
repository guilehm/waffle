package ports

import "tmdb/internal/app/domain"

type MovieService interface {
	MovieByID(id string) (*domain.Movie, error)
	SearchMovies(query string) ([]*domain.Movie, error)
}

package ports

import "tmdb/internal/app/domain"

type MovieRepository interface {
	FindByID(id string) (*domain.Movie, error)
}

package repo

import (
	"tmdb/internal/app/domain"
	"tmdb/internal/app/ports"
)

type MovieRepo struct {
	// db connection
}

func NewMovieRepo() ports.MovieRepository {
	return &MovieRepo{}
}

func (r *MovieRepo) FindByID(id string) (*domain.Movie, error) {
	return nil, nil
}

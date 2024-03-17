package repo

import (
	"tmdb/internal/app/domain"
	"tmdb/internal/app/ports"
)

type MovieRepo struct {
	// db connection
}

func NewMovieRepo( /* db connection */ ) ports.MovieRepository {
	return &MovieRepo{ /* db connection */ }
}

func (r *MovieRepo) FindByID(id string) (*domain.Movie, error) {
	return nil, nil
}

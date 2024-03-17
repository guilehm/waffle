package usecase

import (
	"tmdb/internal/app/domain"
	"tmdb/internal/app/ports"
)

type TMDBUseCase struct {
	tmdbService ports.TMDBService
}

func NewTMDBUseCase(tmdbService ports.TMDBService) *TMDBUseCase {
	return &TMDBUseCase{
		tmdbService: tmdbService,
	}
}

func (uc *TMDBUseCase) MovieDetails(id string) (*domain.Movie, error) {
	return uc.tmdbService.MovieByID(id)
}

func (uc *TMDBUseCase) FindMovies(query string) ([]*domain.Movie, error) {
	return uc.tmdbService.SearchMovies(query)
}

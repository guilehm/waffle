package usecase

import (
	"tmdb/internal/app/domain"
	"tmdb/internal/app/ports"
)

type MovieUseCase struct {
	tmdbService ports.TMDBService
}

func NewMovieUseCase(tmdbService ports.TMDBService) *MovieUseCase {
	return &MovieUseCase{
		tmdbService: tmdbService,
	}
}

func (uc *MovieUseCase) MovieDetails(id string) (*domain.Movie, error) {
	return uc.tmdbService.MovieByID(id)
}

func (uc *MovieUseCase) FindMovies(query string, page int) (*domain.MovieSearchResponse, error) {
	return uc.tmdbService.SearchMovies(query, page)
}

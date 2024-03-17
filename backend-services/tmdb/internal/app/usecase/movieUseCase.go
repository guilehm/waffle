package usecase

import (
	"tmdb/internal/app/domain"
	"tmdb/internal/app/ports"
)

type MovieUseCase struct {
	movieService ports.MovieService
}

func NewMovieUseCase(movieService ports.MovieService) *MovieUseCase {
	return &MovieUseCase{
		movieService: movieService,
	}
}

func (uc *MovieUseCase) MovieDetails(id string) (*domain.Movie, error) {
	return uc.movieService.MovieByID(id)
}

func (uc *MovieUseCase) FindMovies(query string) ([]*domain.Movie, error) {
	return uc.movieService.SearchMovies(query)
}

package api

import (
	"tmdb/internal/app/usecase"
)

type TMDBHandler struct {
	movieUseCase *usecase.MovieUseCase
}

func NewTMDBHandler(movieUseCase *usecase.MovieUseCase) *TMDBHandler {
	return &TMDBHandler{
		movieUseCase: movieUseCase,
	}
}

func (h *TMDBHandler) MovieByID() {
	// call use case and handle responses
}

func (h *TMDBHandler) SearchMovies() {
	// call use case and handle responses
}

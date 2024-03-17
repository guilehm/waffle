package api

import (
	"tmdb/internal/app/usecase"
)

type MovieHandler struct {
	movieUseCase *usecase.MovieUseCase
}

func NewMovieHandler(movieUseCase *usecase.MovieUseCase) *MovieHandler {
	return &MovieHandler{movieUseCase: movieUseCase}
}

func (h *MovieHandler) MovieByID() {
	// call use case and handle responses
}

func (h *MovieHandler) SearchMovies() {
	// call use case and handle responses
}

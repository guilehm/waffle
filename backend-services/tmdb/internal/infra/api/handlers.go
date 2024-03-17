package api

import (
	"tmdb/internal/app/usecase"
)

type TMDBHandler struct {
	tmdbUseCase *usecase.TMDBUseCase
}

func NewTMDBHandler(tmdbUseCase *usecase.TMDBUseCase) *TMDBHandler {
	return &TMDBHandler{tmdbUseCase: tmdbUseCase}
}

func (h *TMDBHandler) MovieByID() {
	// call use case and handle responses
}

func (h *TMDBHandler) SearchMovies() {
	// call use case and handle responses
}

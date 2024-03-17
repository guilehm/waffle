package api

import (
	"tmdb/internal/app/usecase"
)

type TMDBHandler struct {
	movieUseCase *usecase.TMDBUseCase
}

func NewTMDBHandler(movieUseCase *usecase.TMDBUseCase) *TMDBHandler {
	return &TMDBHandler{movieUseCase: movieUseCase}
}

func (h *TMDBHandler) TMDBByID() {
	// call use case and handle responses
}

func (h *TMDBHandler) SearchTMDBs() {
	// call use case and handle responses
}

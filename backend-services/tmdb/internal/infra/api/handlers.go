package api

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"tmdb/internal/app/usecase"
	"tmdb/pkg/errs"
)

type TMDBHandler struct {
	movieUseCase *usecase.MovieUseCase
}

func NewTMDBHandler(movieUseCase *usecase.MovieUseCase) *TMDBHandler {
	return &TMDBHandler{
		movieUseCase: movieUseCase,
	}
}

func (h *TMDBHandler) MovieByID(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")
	movieResponse, err := h.movieUseCase.MovieDetails(movieID)
	if err != nil {
		handleError(w, err)
		return
	}
	handleResponse(w, movieResponse)
}

func (h *TMDBHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	moviesResponse, err := h.movieUseCase.FindMovies(query, page)
	if err != nil {
		handleError(w, err)
		return
	}
	handleResponse(w, moviesResponse)
}

func handleResponse(w http.ResponseWriter, data any) {
	response, err := json.Marshal(data)
	if err != nil {
		handleApiErrors(w, http.StatusInternalServerError, "")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		var apiError *errs.APIError
		switch {
		case errors.As(err, &apiError):
			handleApiErrors(w, apiError.StatusCode, apiError.Message)
		default:
			handleApiErrors(w, http.StatusInternalServerError, "")
		}
		return
	}
}

func handleApiErrors(w http.ResponseWriter, status int, message string) {
	if message == "" {
		message = http.StatusText(status)
	}

	response, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{message})
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

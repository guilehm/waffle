package api

import "github.com/go-chi/chi/v5"

func MoviesRouter(handler *TMDBHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/search", handler.SearchMovies)
		r.Get("/{id}", handler.MovieByID)
	}
}

package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"tmdb/internal/app/usecase"
)

func SetupRouter(movieUseCase *usecase.MovieUseCase) *chi.Mux {
	// create handlers
	tmdbHandler := NewTMDBHandler(movieUseCase)

	// create router
	r := chi.NewRouter()

	// set middlewares
	r.Use(LogRequest)

	// set routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("eat waffles"))
	})
	r.Route("/movies", MoviesRouter(tmdbHandler))

	return r
}

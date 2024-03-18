package ports

import "tmdb/internal/app/domain"

type TMDBService interface {
	MovieService
}

type MovieService interface {
	MovieByID(id string) (*domain.Movie, error)
	SearchMovies(query string, page int) (*domain.MovieSearchResponse, error)
}

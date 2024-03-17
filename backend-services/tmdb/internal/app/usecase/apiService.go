package usecase

import "tmdb/internal/app/domain"

type ApiService struct {
	// repo
}

func NewApiService() *ApiService {
	return &ApiService{}
}

func (s *ApiService) MovieByID(id string) (*domain.Movie, error) {
	return nil, nil
}

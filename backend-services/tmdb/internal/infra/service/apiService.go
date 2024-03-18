package service

import (
	"log"
	"net/http"
	"net/url"
	"time"
	"tmdb/internal/app/ports"
)

type TMDBAPIClient struct {
	baseURL url.URL
	client  *http.Client
	apiKey  string
}

func NewTMDBAPIClient(apiKey string, timeout int) ports.TMDBService {
	baseURL, err := url.Parse("https://api.themoviedb.org/3/")
	if err != nil {
		log.Fatalln("parsing tmdb api base url")
	}

	return &TMDBAPIClient{
		baseURL: *baseURL,
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		apiKey: apiKey,
	}
}

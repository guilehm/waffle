package service

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"tmdb/internal/app/domain"
)

func (c *TMDBAPIClient) MovieByID(id string) (*domain.Movie, error) {
	u := c.MakeURL(EndpointMovies.Detail(id), nil)
	req, err := c.CreateRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	var movie *domain.Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (c *TMDBAPIClient) SearchMovies(query string, page int) (*domain.MovieSearchResponse, error) {
	params := &url.Values{
		"page":  []string{strconv.Itoa(page)},
		"query": []string{url.QueryEscape(query)},
	}
	u := c.MakeURL(EndpointMovies.Search(), params)

	req, err := c.CreateRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	var movieSearchResponse *domain.MovieSearchResponse
	err = json.Unmarshal(body, &movieSearchResponse)
	if err != nil {
		return nil, err
	}
	return movieSearchResponse, nil
}

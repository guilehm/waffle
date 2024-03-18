package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"tmdb/pkg/errs"
)

func (c *TMDBAPIClient) MakeURL(endpoint Endpoint, params *url.Values) string {
	baseURL := c.baseURL
	baseURL.Path = fmt.Sprintf("3/%s", endpoint)
	if params != nil {
		baseURL.RawQuery = params.Encode()
	}
	return baseURL.String()
}

func (c *TMDBAPIClient) CreateRequest(method, url string, data io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, errors.New("creating request")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.apiKey))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	return req, nil
}

func (c *TMDBAPIClient) MakeRequest(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := &errs.APIError{
			StatusCode: resp.StatusCode,
			Message:    http.StatusText(resp.StatusCode),
		}
		return nil, err
	}
	return io.ReadAll(resp.Body)
}

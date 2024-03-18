package service

import "fmt"

type Endpoint string

const (
	EndpointMovies Endpoint = "movie"
)

func (e Endpoint) String() string {
	return string(e)
}

func (e Endpoint) Detail(id string) Endpoint {
	return Endpoint(fmt.Sprintf("%v/%v", e, id))
}

func (e Endpoint) Search() Endpoint {
	return Endpoint(fmt.Sprintf("search/%v", e))
}

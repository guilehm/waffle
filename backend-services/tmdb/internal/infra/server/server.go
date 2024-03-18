package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net"
	"net/http"
	"time"
)

func CreateServer(r *chi.Mux, port string) (*http.Server, *net.Listener, error) {
	addr := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, nil, err
	}
	return server, &listener, nil
}

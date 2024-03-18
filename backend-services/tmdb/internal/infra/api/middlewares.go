package api

import (
	"log/slog"
	"net/http"
	"tmdb/pkg/logging"
)

func LogRequest(next http.Handler) http.Handler {
	logger := logging.Logger
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.InfoContext(
			r.Context(),
			"request received",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("remoteAddr", r.RemoteAddr),
			slog.String("userAgent", r.UserAgent()),
		)
		next.ServeHTTP(w, r)
	})
}

package logging

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	var handler slog.Handler = slog.NewJSONHandler(os.Stdout, opts)
	return slog.New(handler)
}

func init() {
	Logger = NewLogger()
}

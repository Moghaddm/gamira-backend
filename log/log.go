package log

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	env := os.Getenv("APP_ENV")

	if env == "development" {
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

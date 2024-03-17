package logger

import (
	"context"
	"log/slog"
	"os"
)

func Load(_ context.Context) {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, opts))
	slog.SetDefault(logger)
	slog.Debug("logger loaded")
}

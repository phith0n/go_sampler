package logging

import (
	"log/slog"
	"os"

	"go_sampler/providers/config"
)

func NewLogging(cfg *config.Config) *slog.Logger {
	var log *slog.Logger
	options := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	if cfg.Debug {
		options.Level = slog.LevelDebug
		log = slog.New(slog.NewTextHandler(os.Stderr, &options))
	} else {
		log = slog.New(slog.NewJSONHandler(os.Stderr, &options))
	}

	slog.SetDefault(log)
	return log
}

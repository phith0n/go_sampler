package logging

import (
	"go_sampler/providers/config"
	"log/slog"
	"os"
)

func NewLogging(config *config.Config) *slog.Logger {
	var log *slog.Logger
	options := slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}
	if config.Debug {
		options.Level = slog.LevelDebug
		log = slog.New(slog.NewJSONHandler(os.Stderr, &options))
	} else {
		log = slog.New(slog.NewTextHandler(os.Stderr, &options))
	}

	slog.SetDefault(log)
	return log
}

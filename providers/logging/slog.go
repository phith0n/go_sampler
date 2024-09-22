package logging

import (
	"log/slog"
	"os"
	"time"

	"go_sampler/providers/config"

	"github.com/lmittmann/tint"
)

func New(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	if cfg.Debug {
		log = slog.New(tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.RFC3339Nano,
			AddSource:  true,
		}))
	} else {
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		}))
	}

	log.Info("initial slog successfully", "debug", cfg.Debug)
	slog.SetDefault(log)
	return log
}

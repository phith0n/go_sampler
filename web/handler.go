package web

import (
	"log/slog"

	"go_sampler/providers/config"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewHandler(logger *slog.Logger, cfg *config.Config, api *APIHandler) *gin.Engine {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(newLoggerMiddleware(logger), gin.Recovery())
	r.GET("/ping", api.PingHandler)
	return r
}

func newLoggerMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return sloggin.NewWithConfig(logger, sloggin.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelInfo,
		ServerErrorLevel: slog.LevelError,

		WithUserAgent:      true,
		WithRequestID:      true,
		WithRequestBody:    false,
		WithRequestHeader:  false,
		WithResponseBody:   false,
		WithResponseHeader: false,
		WithSpanID:         false,
		WithTraceID:        false,

		Filters: []sloggin.Filter{},
	})
}

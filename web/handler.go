package web

import (
	"log/slog"

	"go_sampler/providers/config"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewHandler(logger *slog.Logger, cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(sloggin.New(logger), gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	return r
}

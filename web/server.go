package web

import (
	"context"
	"net/http"
	"time"

	"go_sampler/providers/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewWebServer(lc fx.Lifecycle, cfg *config.Config, engine *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:              cfg.WebAddr,
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           engine.Handler(),
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				_ = server.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return server
}

package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go_sampler/providers/config"
	"net/http"
)

func NewWebServer(lc fx.Lifecycle, config *config.Config, engine *gin.Engine) *http.Server {
	server := &http.Server{Addr: config.WebAddr, Handler: engine.Handler()}
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

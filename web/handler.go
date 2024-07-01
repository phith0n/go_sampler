package web

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go_sampler/logging"
	"go_sampler/providers/config"
	"time"
)

func NewHandler(config *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(logging.GetLogger(), time.RFC3339, false), gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	return r
}

package web

import (
	"time"
	"go_sampler/logging"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func StartGin(listen string) error {
	r := gin.New()
	r.Use(ginzap.Ginzap(logging.GetLogger(), time.RFC3339, false), gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	return r.Run(listen)
}

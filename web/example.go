package web

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (a *APIHandler) PingHandler(c *gin.Context) {
	slog.Info("example logging")
	c.Data(200, "text/plain", []byte("pong"))
}

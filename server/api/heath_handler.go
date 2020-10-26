package api

import (
	"io"

	"github.com/gin-gonic/gin"
	. "github.com/luckyweiwei/base/logger"
)

func HealthHandler(c *gin.Context) {
	Log.Debug("entering...")

	w := c.Writer

	io.WriteString(w, "health")
}

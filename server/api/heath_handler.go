package api

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"/base/logger"
)

func HealthHandler(c *gin.Context) {
	Log.Debug("entering...")

	w := c.Writer

	io.WriteString(w, "health")
}

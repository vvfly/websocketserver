package api

import (
	"io"

	. "git.nzajiw.com/base/logger"
	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	Log.Debug("entering...")

	w := c.Writer

	io.WriteString(w, "health")
}

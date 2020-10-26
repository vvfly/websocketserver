package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/proto"
)

func WebsocketRequestResolver() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method != http.MethodGet {
			logger.Log.Errorf("Invalid Method. Method=%v", c.Request.Method)
			c.Abort()
			return
		}

		err := ValidCheck(c)
		if err != nil {
			logger.Log.Error(err)
			c.Abort()
			return
		}
	}
}

func ValidCheck(c *gin.Context) error {
	wsConnectInfo := strings.Replace(c.Param("info"), "/", "", 1)
	c.Set(proto.WEBSOCKET_PROTO_REQUEST_CONNECT_INFO, wsConnectInfo)
	return nil
}

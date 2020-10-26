package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
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
	wsCon := proto.NewWsConnectInfo(wsConnectInfo)
	if wsCon == nil {
		logger.Log.Error("wsConnectInfo is nil")
		return errors.New("Encrypt Connect Info error.")
	}

	// check timestamp
	timeNowOfS := utils.GetTimeOfS()
	timestamp := wsCon.Authen.TimeStamp
	if timestamp <= 0 || timeNowOfS-timestamp > 10 {
		logger.Log.Errorf("request timestamp invaild. request timestamp=%v, timestampNow=%v", timestamp, timeNowOfS)
		return errors.New("request timestamp invaild")
	}

	clientIP := c.ClientIP()
	wsCon.Authen.UsrInfo.ClientIP = clientIP
	wsCon.Authen.UsrInfo.LastPingTime = timeNowOfS
	wsCon.Authen.UsrInfo.LastEntryLiveTime = timeNowOfS

	logger.Log.Debugf("Encrypt Connect Info. wsConnectInfo=%v", utils.FormatStruct(wsCon))
	c.Set(proto.WEBSOCKET_PROTO_REQUEST_CONNECT_INFO, wsCon)
	return nil
}

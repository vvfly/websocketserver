package server

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/luckyweiwei/base/grmon"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/server/api"
	"github.com/luckyweiwei/websocketserver/server/middleware"
	"github.com/luckyweiwei/websocketserver/server/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "github.com/luckyweiwei/base/logger"
)

const (
	WEBSOCKET_CONTEXT      = "/websocket"
	WEBSOCKET_SESSION_NAME = "/*info"
)

func ServerStart() {
	debugMode := model.GetServerConfig().DebugMode
	if debugMode != 1 {
		gin.SetMode(gin.ReleaseMode)
	}

	grm := grmon.GetGRMon()
	grm.Go("websocketServer", websocketServer)
	grm.Go("apiServer", apiServer)
}

func websocketServer() {

	port := model.GetServerConfig().WebsocketPort

	defer utils.CatchException()

	router := createDefaultRouter()

	// 设置路由处理函数
	router.GET("/health", api.HealthHandler)

	wsRouter := router.Group(WEBSOCKET_CONTEXT)
	{
		wsRouter.GET(WEBSOCKET_SESSION_NAME, middleware.WebsocketRequestResolver(), websocket.WebsocketHandler)
	}

	runServer(port, router)
}

func apiServer() {

	port := model.GetServerConfig().ApiServerConf.ApiPort

	defer utils.CatchException()

	router := createDefaultRouter()

	// 设置路由处理函数
	router.GET("/health", api.HealthHandler)

	runServer(port, router)
}

func runServer(port int, router *gin.Engine) {
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	Log.Infof("listening at port %v", port)
	err := server.ListenAndServe()
	if err != nil {
		Log.Error(err)
		os.Exit(4)
	}
}

func createDefaultRouter() *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Set-Cookie", "Content-Encoding"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	return router
}

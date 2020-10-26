package server

import (
	"time"

	_ "github.com/luckyweiwei/base/daemon"
	"github.com/luckyweiwei/websocketserver/model"

	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
)

const (
	SERVER_NAME   = "WebsocketServer"
	SERVER_CONFIG = "config/websocket_server.toml"
)

func ServerEntry() {
	// 初始化配置文件
	model.ServerConfigInit()

	// 初始化日志
	serverConfig := model.GetServerConfig()
	LogInitWithConfig(&serverConfig.LogConf)

	utils.UseMaxCpu()

	Log.Infof("#############################################")
	Log.Infof("################ " + SERVER_NAME + " #################")
	Log.Infof("#############################################")

	Log.Info("server init...")
	ServerInit()
	Log.Info("server init done...")

	// 启动 pprof 性能分析端口
	Log.Infof("start http pprof......")
	HttpPprof()

	Log.Info("server start...")
	ServerStart()
	Log.Info("server is running...")

	Log.Infof("keep main func here ... ...")
	for {
		time.Sleep(time.Duration(3) * time.Second)
	}
}

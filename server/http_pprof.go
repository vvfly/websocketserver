package server

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/luckyweiwei/websocketserver/model"
)

func HttpPprof() {
	pprofConfig := model.GetServerConfig().PprofConf

	pport := pprofConfig.PprofPort
	disabled := pprofConfig.Disabled

	if disabled == 1 { // 没配置则不启动
		return
	}

	go func() {
		log.Println(http.ListenAndServe(":"+strconv.Itoa(pport), nil))
	}()
}

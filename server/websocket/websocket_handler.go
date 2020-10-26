package websocket

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/luckyweiwei/websocketserver/helper"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/biz"
	"github.com/luckyweiwei/websocketserver/server/connectionm"

	"github.com/luckyweiwei/base/grmon"
	"github.com/luckyweiwei/base/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "github.com/luckyweiwei/base/logger"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	WriteBufferPool:  &sync.Pool{},
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理 websocket 请求
func WebsocketHandler(c *gin.Context) {
	defer utils.CatchException()

	Log.Debug("entering...")

	connInfo, ok := c.Get(proto.WEBSOCKET_PROTO_REQUEST_CONNECT_INFO)
	if !ok {
		Log.Error("can't get connect info")
		return
	}
	wsConnInfo, ok := connInfo.(*proto.WsConnectInfo)
	if !ok {
		Log.Errorf("valid wsConnInfo.wsConnInfo=%v", wsConnInfo)
		return
	}
	sessionID := wsConnInfo.SessionsInfoStr

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		Log.Errorf("Failed to upgrade http to websocket, err = %v", err)
		return
	}

	// 更新用户在线信息
	proto.UserOnlineInfoConnChan <- wsConnInfo

	serverConfig := model.GetServerConfig()

	// 加入connection manager 管理
	cm := connectionm.GetConnectionManager()
	connection := cm.NewConnection(conn, wsConnInfo)
	dataCh := connection.DataCh

	// 进入直播间发入场消息
	grm := grmon.GetGRMon()
	grm.Go(sessionID+"-Entry", func() {
		biz.EntryHandler(connection.WsConnectInfo)
	})

	readData := func() error {
		_, data, err := conn.ReadMessage()
		if err != nil {
			// 检测到用户主动断开连接
			Log.Warn(err)

			biz.DisconnectHandler(proto.MT_DISCONNECT, sessionID, "")

			return err
		}

		// ping 消息处理
		if string(data) == proto.WEBSOCKET_PING {
			// Log.Debugf("len = %v, read reqData = \n%v", len(data), string(data))

			// 更新最后心跳时间
			connection.WsConnectInfo.Authen.UsrInfo.LastPingTime = utils.GetTimeOfS()

			// 更新用户在线信息
			proto.UserOnlineInfoMsgChan <- wsConnInfo

			return nil
		}

		reqData, err := helper.Des3CBCDecrypt4WebsocketMsg([]byte(serverConfig.Des3Key4WsMsg), data)
		if err != nil {
			Log.Warn(err)
			return err
		}

		// 老版本 ping消息加密处理
		if string(reqData) == proto.WEBSOCKET_PING {
			// Log.Debugf("len = %v, read reqData = \n%v", len(data), string(data))

			// 更新最后心跳时间
			connection.WsConnectInfo.Authen.UsrInfo.LastPingTime = utils.GetTimeOfS()

			// 更新用户在线信息
			proto.UserOnlineInfoMsgChan <- wsConnInfo

			return nil
		}

		Log.Infof("read reqData =%v, len = %v", string(reqData), len(reqData))

		grm.Go(sessionID+"-Msg", func() {
			biz.BizHandler(sessionID, string(reqData))
		})

		return nil
	}

	writeData := func() error {
		data, ok := <-dataCh
		if !ok {
			emsg := "ch is closed"
			Log.Warnf(emsg)
			return errors.New(emsg)
		}

		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			Log.Warn(err)
		}

		Log.Debugf("Finish write data for sessionID = %v", sessionID)

		return nil
	}

	reader := func() {
		Log.Debugf("reader for user = %v", sessionID)
		for {
			err := readData()
			if err != nil {
				cm.ClearConnection(connection)
				return
			}
		}
	}

	// reader 负责关闭链接, writer只需退出线程
	writer := func() {
		Log.Debugf("writer for user = %v", sessionID)
		for {
			err := writeData()
			if err != nil {
				return
			}
		}
	}

	readerName := fmt.Sprintf("reader_%v", sessionID)
	writerName := fmt.Sprintf("writer_%v", sessionID)

	grm.Go(readerName, reader)
	grm.Go(writerName, writer)
}

package broadcast

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/luckyweiwei/base/grmon"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
)

const (
	GroupLen = 1000
)

func GroupWriteData(conns []*connectionm.Connection, data []byte) {

	grm := grmon.GetGRMon()

	lenConns := len(conns)
	groups := lenConns / GroupLen

	if groups <= 0 {
		grm.Go("GroupWriteData", func() {
			WriteData(conns, data)
		})
		return
	}

	for i := 0; i < groups+1; i++ {
		index := i
		if (index+1)*GroupLen >= lenConns {
			grm.Go("GroupWriteData", func() {
				WriteData(conns[index*GroupLen:], data)
			})
			return
		}

		grm.Go("GroupWriteData", func() {
			WriteData(conns[index*GroupLen:(index+1)*GroupLen], data)
		})
	}
}

func WriteData(conns []*connectionm.Connection, data []byte) {
	start := time.Now()
	for _, connection := range conns {
		if connection == nil {
			continue
		}

		conn := connection.Conn
		if conn == nil {
			continue
		}

		conn.WriteMessage(websocket.BinaryMessage, data)
	}

	// debug
	cost := time.Since(start)
	if cost > time.Second {
		Log.Infof("oh my god. cost=%v", cost)
	}
}

func WriteDataToCh(conns []*connectionm.Connection, data []byte) {
	connNexts := make([]*connectionm.Connection, 0)

	for _, connection := range conns {
		dataCh := connection.DataCh
		lenDataCh := len(dataCh)
		if lenDataCh == connectionm.DATA_CHAN_MAX_LEN {
			connNexts = append(connNexts, connection) // 如果连接通道已满，累积到下次发送
			continue
		}

		safeSend := func() { // 小概率发送到已关闭的通道，简单做recover处理
			defer func() {
				err := recover()
				if err != nil {
					Log.Errorf("panic captured, ch is closed, err = %v", err)
				}
			}()

			dataCh <- data
		}
		safeSend()
		time.Sleep(time.Nanosecond)
	}

	if len(connNexts) > 0 {
		NextWriteDataToCh(connNexts, data)
	}
}

func NextWriteDataToCh(conns []*connectionm.Connection, data []byte) {

	Log.Errorf("bad network. connection data channel is full! lenConn=%v", len(conns))
	time.Sleep(time.Second)

	for _, connection := range conns {
		dataCh := connection.DataCh
		lenDataCh := len(dataCh)
		if lenDataCh == connectionm.DATA_CHAN_MAX_LEN { // 如果连接通道已满，消息丢掉
			continue
		}

		safeSend := func() { // 小概率发送到已关闭的通道，简单做recover处理
			defer recover()
			dataCh <- data
		}
		safeSend()
		time.Sleep(time.Nanosecond)
	}
}

package wstest

import (
	"testing"
	"time"

	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/helper"

	"github.com/gorilla/websocket"
)

var (
	// WebsocketAddr = `ws://localhost:22000/websocket/vlZQuRVwEL/lzYGgPPEirIUwPDHnF7qNrfFSTq6T4SVuolF2u2IK50toqUafNlDPYLShPIjM3h2StCRGehCxAppWOXPK0Mpy9sjA6nLNl+jSFqZGycvUyKQ9sTc5H/XaGhYcfZvQRmIZn4yiEZnQFooy104KdDYcybSb/bSOJ5Ue9Mv79HSOhEr5dTII0JEKYbjufrltlxXr/rzkcNnKy/rHT59kPWe7wgxxtEvzxpxkEk9AABYTCjeXRC4VHgCvmrmnKvr6UQzZ3lt5xCc1diLmdjYTIJ9k+xfXR7A4JtXqe+ci/ddIgcTPyktX7C2uNRVscCt0XdqbIjVElw2kr7QsDorLCatGPos/Jr12WEgm5GbREi94sY+8V2wQdho9U3PRAYNRhku3Yb17dGhhuBe5fKDL3+s9pZH9QPTbxrL3J8yYKmuyVplO8q8unn2N2A2XrEeMKmpSL+ZtEH+VSkYYGFttO4nl1rX7BQWxN9j7DAX7B9QbMQy2fWoDsEb2c18zmI9ZLAEQ/Q6yCU9Kb4jkrNmArHjTz0THXz/DByr5HdDxLkxiar1X6ZZPmMbaHXNMnnx4UPqisFiEJDJLmz7Ae4JQtP38Qp8zeWeErMqbn7wnGmdn1eygoP69B4rUGpIhOQXefm1aD41PLNncY8ztjEqdVr7OlIjHdV9krsshAeivlzbyjJHcb25S37AdGaU5VdIecloAH0yicif876iZ/Hch5gTglm/NUUVnfWnNZxj0OSSwFCpNMEaBDMj/kwBtwbUBOQAM1TZEWySS0sc5P3EH8sCFtaC7r0D722kTGgql7ZIBrEJ2GgmAKRnUwEP5d5NgOaKT7rMnVQWS21RrfhcVmb4xwSUPB9qKVJk2266FSQJWRFHZQyY+iPhlHT5Vw5PQGJT4DqeQn50b1uGBGHzxL02rCLN4RE9MQeoX2cdEAKXsGAf4kr/17lM9iP88Dw3r+Z4h53R+5ymasl73N1jRafgXYCg3Fu9Ua49taVL0TUAbr7QVxG+Oka5cVPEO3sOf3h+MqRsbjvKn1H2wkZWbMkgfpRp7xd+jVRxAeQniJuDVrVZ7D0fYbayqE4mPmkfB/dBSI0OAkRqMDVQfzI7KZb8J9c8onnrv3AJD2ktPs1l78bByF8Fgu40Q7AXWpZMFFKaFh/gsUkBsGZZw6LEWRunZ+glC3RFpfUXG3BHmgBZ+bMKHwTmDIWkUHsWxa/+xyvsAdJT82g3DLkq+KT19cYK8UqiMZpnUDbU9V1uBINQYEn/kAkEvjHjx9yHPaUFzK4ua9QrVf4GVAO1SN9agccutaOCpiU7zVG+RR5RtvoBaGi82CYjS60EvEG+a+1bIs8LGhDzzfad+9z2QzksiD9cZSApFYFroxIN8i2bWumBTyKRUx5EpSOMvkE+cr2F+aJekpd63NMPngTc4kklCdo7QhYsfWr64GyDSmrOzxPslWKMj5WBoH3q8mLJiQ6yg1o1da7vIn1r8KWIXYBscsxUq4Ib/g1z29ycehToGAMb9tg8PF3Sn608/6zbyu7oQiu2su4KqDdg6YqMZQCfzfubRuVhCotWpL02tLxqj1jbYp3MDm5/K1GEvDYDbvmQOuqlLHwU22FncEe7tyLZb1dAgSm5HC4GMi7Czn5BGRaDvWg==`
	WebsocketAddr = `ws://localhost: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`
	MsgPing       = "ping"

	MsgChat = `
	{
		"businessData": {
		  "content": "甚么",
		  "openDanmu": "0",
		  "openNobilityDanmu": "0"
		},
		"messageType": "chat",
		"r": "mdMNltE3slB6wb49",
		"s": "e461be961aa94a9cacd699f2cd7a10b6",
		"t": "1593325916"
	  }
	`

	MsgChatReceipt = `
	{
		"messageType": "chatReceipt",
		"businessData": {
			"senderId": "发送人ID（广播聊天的userId）",
			"messageId": "202cb962ac59075b964b07152d234b70",
			"status": "1"
		},
		"r": "随机字符串",
		"t": "时间戳",
		"s": "签名 md5(businessData值+t+r+秘钥)"
	}
	`
)

func TestMsgChatReceipt(t *testing.T) {
	c, _, err := websocket.DefaultDialer.Dial(WebsocketAddr, nil)
	if err != nil {
		Log.Error("err = %v", err)
	}
	defer c.Close()

	sendAndRecv(c, MsgChatReceipt)
}

func TestMsgChat(t *testing.T) {

	c, _, err := websocket.DefaultDialer.Dial(WebsocketAddr, nil)
	if err != nil {
		Log.Error("err = %v", err)
	}
	defer c.Close()

	time.Sleep(3 * time.Second)

	sendAndRecv(c, MsgChat)

	time.Sleep(3 * time.Second)
}

func TestMsgPing(t *testing.T) {

	c, _, err := websocket.DefaultDialer.Dial(WebsocketAddr, nil)
	if err != nil {
		Log.Error("err = %v", err)
	}
	defer c.Close()

	count := 0
	for {
		count++
		if count > 5 {
			break
		}

		err := c.WriteMessage(websocket.BinaryMessage, []byte(MsgPing))
		if err != nil {
			Log.Error("err = %v", err)
			return
		}

		_, message, err := c.ReadMessage()
		if err != nil {
			Log.Error("err = %v", err)
			return
		}
		Log.Infof("resp msg = %v", string(message))

		time.Sleep(3 * time.Second)
	}
}

func sendAndRecv(c *websocket.Conn, msg string) string {
	data, err := helper.Des3CBCEncrypt4WebsocketMsg([]byte("c21d31be-4300-4881-a553-156ebb5df087"), []byte(msg))
	if err != nil {
		Log.Error(err)
	}

	err = c.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		Log.Error(err)
	}

	_, message, err := c.ReadMessage()
	if err != nil {
		Log.Error("err = %v", err)
		return ""
	}

	// 解密
	respData, err := helper.Des3CBCDecrypt4WebsocketMsg([]byte("c21d31be-4300-4881-a553-156ebb5df087"), message)
	if err != nil {
		Log.Error("err = %v", err)
		return ""
	}

	respDataStr := string(respData)
	Log.Debugf("resp message=%v", utils.FormatJsonStr(respDataStr))

	return respDataStr
}

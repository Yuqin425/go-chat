package chat

import (
	"IM_QQ/settings"
	"IM_QQ/utils"
	"encoding/json"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Unregister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			zap.L().Error("client read message error", zap.Any("client read message error", err.Error()))
			MyServer.Unregister <- c
			c.Conn.Close()
			break
		}

		msg := &utils.Message{}
		json.Unmarshal(message, msg)

		// pong
		if msg.Type == utils.HEART_BEAT {
			pong := &utils.Message{
				Content: utils.PONG,
				Type:    utils.HEART_BEAT,
			}
			pongByte, err2 := json.Marshal(pong)
			if nil != err2 {
				zap.L().Error("client marshal message error", zap.Any("client marshal message error", err2.Error()))
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			if settings.Conf.ChannelType == utils.NSQ {
				Send(message)
			} else {
				MyServer.Broadcast <- message
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}

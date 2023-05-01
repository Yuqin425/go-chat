package router

import (
	"IM_QQ/chat"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocekt(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		return
	}
	zap.L().Info("newUser", zap.String("newUser", user))
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &chat.Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}

	chat.MyServer.Register <- client
	go client.Read()
	go client.Write()
}

package router

import (
	"IM_QQ/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	group := r.Group("")
	{
		group.GET("/user", api.GetUserList)
		group.GET("/user/details", api.GetUserDetails)
		group.GET("/user/name", api.GetUserOrGroupByName)
		group.POST("/user/register", api.Register)
		group.POST("/user/login", api.Login)
		group.PUT("/user", api.ModifyUserInfo)

		group.POST("/friend", api.AddFriend)
		group.POST("/friendDelete", api.DeleteFriend)
		group.GET("/message", api.GetMessage)
		group.POST("/file", api.SaveFile)

		group.GET("/group/:uuid", api.GetGroup)
		group.POST("/group/:uuid", api.SaveGroup)
		group.POST("/group/join/:userUuid/:groupUuid", api.JoinGroup)
		group.GET("/group/user/:uuid", api.GetGroupUsers)

		group.GET("/chat", RunSocekt)
	}

	return r
}

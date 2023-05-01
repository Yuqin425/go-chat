package api

import (
	"IM_QQ/service"
	"IM_QQ/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取消息列表
func GetMessage(c *gin.Context) {
	zap.L().Info(c.Query("uuid"))
	var messageRequest utils.MessageRequest
	err := c.ShouldBindJSON(&messageRequest)
	if err != nil {
		zap.L().Error("bindQueryError", zap.Any("bindQueryError", err))
		return
	}
	zap.L().Info("messageRequest params: ", zap.Any("messageRequest", messageRequest))

	messages, err := service.MessageService.GetMessages(messageRequest)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
		return
	}

	utils.ResponseSuccess(c, messages)
}

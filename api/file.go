package api

import (
	"IM_QQ/service"
	"IM_QQ/settings"
	"IM_QQ/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strings"
)

// 上传头像等文件
func SaveFile(c *gin.Context) {
	namePreffix := uuid.New().String()

	userUuid := c.PostForm("uuid")

	file, _ := c.FormFile("file")
	fileName := file.Filename
	index := strings.LastIndex(fileName, ".")
	suffix := fileName[index:]

	newFileName := namePreffix + suffix

	zap.L().Info("file", zap.Any("file name", settings.Conf.FilePath+newFileName))
	zap.L().Info("userUuid", zap.Any("userUuid name", userUuid))

	c.SaveUploadedFile(file, settings.Conf.FilePath+newFileName)
	err := service.UserService.ModifyUserAvatar(newFileName, userUuid)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
		return
	}
	utils.ResponseSuccess(c, newFileName)
}

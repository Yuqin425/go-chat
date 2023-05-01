package api

import (
	"IM_QQ/models"
	"IM_QQ/service"
	"IM_QQ/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, "缺少必填字段")
		return
	}
	confirmPassword := c.PostForm("confirm_password")
	err = service.UserService.Register(user.Username, user.Password, confirmPassword)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
		return
	}

	utils.ResponseSuccess(c, "注册成功")
}

func Login(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	zap.L().Debug("user", zap.Any("user", user))

	if err := service.UserService.Login(&user); err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
		return
	}

	utils.ResponseSuccess(c, "登陆成功")
}

func ModifyUserInfo(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	zap.L().Debug("user", zap.Any("user", user))
	if err := service.UserService.ModifyUserInfo(&user); err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
		return
	}

	utils.ResponseSuccess(c, "修改成功")
}

func GetUserDetails(c *gin.Context) {
	uuid := c.Query("uuid")

	utils.ResponseSuccess(c, service.UserService.GetUserDetails(uuid))
}

// 通过用户名获取用户信息
func GetUserOrGroupByName(c *gin.Context) {
	name := c.Query("name")

	utils.ResponseSuccess(c, service.UserService.GetUserOrGroupByName(name))
}

func GetUserList(c *gin.Context) {
	uuid := c.Query("uuid")
	data, err := service.UserService.GetUserList(uuid)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
	}
	utils.ResponseSuccess(c, data)
}

func AddFriend(c *gin.Context) {
	var userFriendRequest utils.FriendRequest
	c.ShouldBindJSON(&userFriendRequest)

	err := service.UserService.AddFriend(&userFriendRequest)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
		return
	}

	utils.ResponseSuccess(c, "添加成功")
}

func DeleteFriend(c *gin.Context) {
	var userFriendRequest utils.FriendRequest
	c.ShouldBindJSON(&userFriendRequest)

	err := service.UserService.DeleteFriend(&userFriendRequest)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
		return
	}

	utils.ResponseSuccess(c, "删除成功")
}

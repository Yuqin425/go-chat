package api

import (
	"IM_QQ/models"
	"IM_QQ/service"
	"IM_QQ/utils"
	"github.com/gin-gonic/gin"
)

// 获取分组
func GetGroup(c *gin.Context) {
	uuid := c.Param("uuid")
	groups, err := service.GroupService.GetGroups(uuid)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}

	utils.ResponseSuccess(c, groups)
}

// 创建
func SaveGroup(c *gin.Context) {
	uuid := c.Param("uuid")
	var group models.Group
	c.ShouldBindJSON(&group)

	err := service.GroupService.SaveGroup(uuid, group)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err.Error())
		return
	}
	utils.ResponseSuccess(c, "保存成功")
}

// 加入组别
func JoinGroup(c *gin.Context) {
	userUuid := c.Param("userUuid")
	groupUuid := c.Param("groupUuid")
	err := service.GroupService.JoinGroup(groupUuid, userUuid)
	if err != nil {
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}
	utils.ResponseSuccess(c, "加入成功")
}

// 获取组内成员信息
func GetGroupUsers(c *gin.Context) {
	groupUuid := c.Param("uuid")
	users := service.GroupService.GetUserIdByGroupUuid(groupUuid)
	utils.ResponseSuccess(c, users)
}

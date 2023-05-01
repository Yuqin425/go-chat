package service

import (
	"IM_QQ/dao"
	"IM_QQ/models"
	"IM_QQ/utils"
	"github.com/google/uuid"
)

type groupService struct {
}

var GroupService = new(groupService)

func (g *groupService) GetGroups(uuid string) ([]utils.GroupResponse, error) {
	db := dao.GetDB()
	var queryUser *models.User
	db.First(&queryUser, "uuid = ?", uuid)

	if queryUser.Id <= 0 {
		return nil, utils.ErrorUserNotExit
	}

	var groups []utils.GroupResponse

	db.Raw("SELECT g.id AS group_id, g.uuid, g.created_at, g.name, g.notice FROM group_members AS gm LEFT JOIN `groups` AS g ON gm.group_id = g.id WHERE gm.user_id = ?",
		queryUser.Id).Scan(&groups)

	return groups, nil
}

func (g *groupService) SaveGroup(userUuid string, group models.Group) error {
	db := dao.GetDB()
	var fromUser models.User
	db.Find(&fromUser, "uuid = ?", userUuid)
	if fromUser.Id <= 0 {
		return utils.ErrorInvalidID
	}

	group.UserId = fromUser.Id
	group.Uuid = uuid.New().String()
	db.Save(&group)

	groupMember := models.GroupMember{
		UserId:   fromUser.Id,
		GroupId:  group.ID,
		Nickname: fromUser.Username,
		Mute:     0,
	}
	db.Save(&groupMember)
	return nil
}

func (g *groupService) GetUserIdByGroupUuid(groupUuid string) []models.User {
	var group models.Group
	db := dao.GetDB()
	db.First(&group, "uuid = ?", groupUuid)
	if group.ID <= 0 {
		return nil
	}

	var users []models.User
	db.Raw("SELECT u.uuid, u.avatar, u.username FROM `groups` AS g JOIN group_members AS gm ON gm.group_id = g.id JOIN users AS u ON u.id = gm.user_id WHERE g.id = ?",
		group.ID).Scan(&users)
	return users
}

func (g *groupService) JoinGroup(groupUuid, userUuid string) error {
	var user models.User
	db := dao.GetDB()
	db.First(&user, "uuid = ?", userUuid)
	if user.Id <= 0 {
		return utils.ErrorUserNotExit
	}

	var group models.Group
	db.First(&group, "uuid = ?", groupUuid)
	if user.Id <= 0 {
		return utils.ErrorInvalidID
	}
	var groupMember models.GroupMember
	db.First(&groupMember, "user_id = ? and group_id = ?", user.Id, group.ID)
	if groupMember.ID > 0 {
		return utils.ErrorInGroup
	}
	nickname := user.Nickname
	if nickname == "" {
		nickname = user.Username
	}
	groupMemberInsert := models.GroupMember{
		UserId:   user.Id,
		GroupId:  group.ID,
		Nickname: nickname,
		Mute:     0,
	}
	db.Save(&groupMemberInsert)

	return nil
}

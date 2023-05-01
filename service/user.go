package service

import (
	"IM_QQ/dao"
	"IM_QQ/models"
	"IM_QQ/utils"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type userService struct {
}

var UserService = new(userService)

func (u *userService) Register(username, password, confirmPassword string) error {
	ok := dao.FindUser(username)
	if ok {
		// fmt.Println(ok)
		return utils.ErrorUserExit
	}
	// fmt.Println(1234)
	user := models.User{}
	user.Username = username

	if user.Password != confirmPassword {
		return utils.ErrorPwdDifferent
	}
	user.Password = utils.MakePassword(password)
	user.Uuid = uuid.New().String()
	user.CreateAt = time.Now()
	user.DeleteAt = 0

	dao.CreateUser(&user)
	return nil
}

func (u *userService) Login(user *models.User) error {
	ok := dao.FindUser(user.Username)
	if !ok {
		// fmt.Println(ok)
		return utils.ErrorUserNotExit
	}
	if ok = dao.CheckPwd(user.Username, user.Password); !ok {
		return utils.ErrorPasswordWrong
	}
	return nil
}

func (u *userService) GetUserDetails(uuid string) models.User {
	var queryUser *models.User
	db := dao.GetDB()
	db.Select("uuid", "username", "nickname", "avatar").First(&queryUser, "uuid = ?", uuid)
	//fmt.Println(queryUser)
	//fmt.Println(123)
	return *queryUser
}

func (u *userService) ModifyUserInfo(user *models.User) error {
	var queryUser *models.User
	db := dao.GetDB()
	db.First(&queryUser, "username = ?", user.Username)
	zap.L().Debug("queryUser", zap.Any("queryUser", queryUser))
	var nullId int32 = 0
	if nullId == queryUser.Id {
		return utils.ErrorUserNotExit
	}
	queryUser.Nickname = user.Nickname
	queryUser.Email = user.Email
	queryUser.Password = user.Password

	db.Save(queryUser)
	return nil
}

func (u *userService) GetUserOrGroupByName(name string) utils.SearchResponse {
	var queryUser *models.User
	db := dao.GetDB()
	db.Select("uuid", "username", "nickname", "avatar").First(&queryUser, "username = ?", name)

	var queryGroup *models.Group
	db.Select("uuid", "name").First(&queryGroup, "name = ?", name)

	search := utils.SearchResponse{
		User:  *queryUser,
		Group: *queryGroup,
	}
	fmt.Println(search)

	return search
}
func (u *userService) GetUserList(uuid string) ([]models.User, error) {
	db := dao.GetDB()

	var queryUser *models.User
	db.First(&queryUser, "uuid = ?", uuid)
	var nullId int32 = 0
	if nullId == queryUser.Id {
		return nil, utils.ErrorUserNotExit
	}

	var queryUsers []models.User
	db.Raw("SELECT u.username, u.uuid, u.avatar FROM user_friends AS uf JOIN users AS u ON uf.friend_id = u.id WHERE uf.user_id = ?", queryUser.Id).Scan(&queryUsers)

	fmt.Println(queryUsers)
	return queryUsers, nil
}
func (u *userService) AddFriend(userFriendRequest *utils.FriendRequest) error {
	var queryUser *models.User
	db := dao.GetDB()
	db.First(&queryUser, "uuid = ?", userFriendRequest.Uuid)
	// fmt.Println(queryUser)
	zap.L().Debug("queryUser", zap.Any("queryUser", queryUser))
	var nullId int32 = 0
	if nullId == queryUser.Id {
		return utils.ErrorUserNotExit
	}

	var friend *models.User
	db.First(&friend, "username = ?", userFriendRequest.FriendUsername)
	if nullId == friend.Id {
		return utils.ErrorFriendUserNotExit
	}

	userFriend := models.UserFriend{
		UserId:   queryUser.Id,
		FriendId: friend.Id,
	}

	var userFriendQuery *models.UserFriend
	db.First(&userFriendQuery, "user_id = ? and friend_id = ?", queryUser.Id, friend.Id)
	// fmt.Println(321)
	if userFriendQuery.ID != nullId {
		return utils.ErrorFriendExit
	}
	db.First(&userFriendQuery, "user_id = ? and friend_id = ?", friend.Id, queryUser.Id)
	if userFriendQuery.ID != nullId {
		return utils.ErrorFriendExit
	}

	db.Save(&userFriend)
	zap.L().Debug("userFriend", zap.Any("userFriend", userFriend))

	return nil
}

func (u *userService) DeleteFriend(userFriendRequest *utils.FriendRequest) error {
	var queryUser *models.User
	db := dao.GetDB()
	db.First(&queryUser, "uuid = ?", userFriendRequest.Uuid)
	// fmt.Println(queryUser)
	zap.L().Debug("queryUser", zap.Any("queryUser", queryUser))
	var nullId int32 = 0
	if nullId == queryUser.Id {
		return utils.ErrorUserNotExit
	}

	var friend *models.User
	db.First(&friend, "username = ?", userFriendRequest.FriendUsername)
	if nullId == friend.Id {
		return utils.ErrorFriendUserNotExit
	}
	var userFriendQuery *models.UserFriend
	db.First(&userFriendQuery, "user_id = ? and friend_id = ?", queryUser.Id, friend.Id)
	// fmt.Println(321)
	if userFriendQuery.ID != nullId {
		db.Delete(&userFriendQuery)
		return nil
	}
	db.First(&userFriendQuery, "user_id = ? and friend_id = ?", friend.Id, queryUser.Id)
	if userFriendQuery.ID != nullId {
		db.Delete(&userFriendQuery)
		return nil
	}

	return utils.ErrorsFriendNotExit
}

// 修改头像
func (u *userService) ModifyUserAvatar(avatar string, userUuid string) error {
	var queryUser *models.User
	db := dao.GetDB()
	db.First(&queryUser, "uuid = ?", userUuid)

	if NULL_ID == queryUser.Id {
		return utils.ErrorUserNotExit
	}

	db.Model(&queryUser).Update("avatar", avatar)
	return nil
}

package service

import (
	"IM_QQ/dao"
	"IM_QQ/models"
	"IM_QQ/utils"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const NULL_ID int32 = 0

type messageService struct {
}

var MessageService = new(messageService)

func (m *messageService) GetMessages(message utils.MessageRequest) ([]utils.MessageResponse, error) {
	db := dao.GetDB()

	if message.MessageType == utils.MESSAGE_TYPE_USER {
		var queryUser *models.User
		db.First(&queryUser, "uuid = ?", message.Uuid)

		if queryUser.Id == NULL_ID {
			return nil, utils.ErrorUserNotExit
		}

		var friend *models.User
		db.First(&friend, "username = ?", message.FriendUsername)
		if friend.Id == NULL_ID {
			return nil, utils.ErrorUserNotExit
		}

		var messages []utils.MessageResponse

		db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar, to_user.username AS to_username  FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id LEFT JOIN users AS to_user ON m.to_user_id = to_user.id WHERE from_user_id IN (?, ?) AND to_user_id IN (?, ?)",
			queryUser.Id, friend.Id, queryUser.Id, friend.Id).Scan(&messages)

		return messages, nil
	}

	if message.MessageType == utils.MESSAGE_TYPE_GROUP {
		messages, err := fetchGroupMessage(db, message.Uuid)
		if err != nil {
			return nil, err
		}

		return messages, nil
	}

	return nil, errors.New("不支持查询类型")
}

func fetchGroupMessage(db *gorm.DB, toUuid string) ([]utils.MessageResponse, error) {
	var group models.Group
	db.First(&group, "uuid = ?", toUuid)
	if group.ID <= 0 {
		return nil, utils.ErrorGroupNotExit
	}

	var messages []utils.MessageResponse

	db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id WHERE m.message_type = 2 AND m.to_user_id = ?",
		group.ID).Scan(&messages)

	return messages, nil
}

func (m *messageService) SaveMessage(message utils.Message) {
	db := dao.GetDB()
	var fromUser models.User
	db.Find(&fromUser, "uuid = ?", message.From)
	if fromUser.Id == NULL_ID {
		zap.L().Error("SaveMessage not find from user", zap.Any("SaveMessage not find from user", fromUser.Id))
		return
	}

	var toUserId int32 = 0

	if message.MessageType == utils.MESSAGE_TYPE_USER {
		var toUser models.User
		db.Find(&toUser, "uuid = ?", message.To)
		if toUser.Id == NULL_ID {
			return
		}
		toUserId = toUser.Id
	}

	if message.MessageType == utils.MESSAGE_TYPE_GROUP {
		var group models.Group
		db.Find(&group, "uuid = ?", message.To)
		if NULL_ID == group.ID {
			return
		}
		toUserId = group.ID
	}

	saveMessage := models.Message{
		FromUserId:  fromUser.Id,
		ToUserId:    toUserId,
		Content:     message.Content,
		ContentType: int16(message.ContentType),
		MessageType: int16(message.MessageType),
		Url:         message.Url,
	}
	db.Save(&saveMessage)
}

package chat

import (
	"IM_QQ/service"
	"IM_QQ/settings"
	"IM_QQ/utils"
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
	"strings"
	"sync"
)

var MyServer = NewServer()

type Server struct {
	Clients    map[string]*Client
	mutex      *sync.Mutex
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:      &sync.Mutex{},
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// 消费nsq里面的消息, 然后直接放入go channel中统一进行消费
func ConsumerNsqMsg(data []byte) {
	MyServer.Broadcast <- data
}

func (s *Server) Start() {
	zap.L().Info("start server", zap.Any("start server", "start server..."))
	for {
		select {
		case conn := <-s.Register:
			zap.L().Info("login", zap.Any("login", "new user login in "+conn.Name))
			s.Clients[conn.Name] = conn
			msg := &utils.Message{
				From:    "System",
				To:      conn.Name,
				Content: "welcome!",
			}
			protoMsg, _ := json.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.Unregister:
			zap.L().Info("loginout", zap.Any("loginout", conn.Name))
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}

		case message := <-s.Broadcast:
			msg := &utils.Message{}
			json.Unmarshal(message, msg)

			//msg.From = "ce7c6b07-105c-492d-a572-3705d5870c3f"
			//msg.To = "b4652631-bbaf-42d7-bb22-e2c9bc61fec2"
			//msg.Content = "1234"
			//msg.ContentType = 1
			//msg.MessageType = 1

			if msg.To != "" {
				// 一般消息，比如文本消息，视频文件消息等
				if msg.ContentType >= utils.TEXT && msg.ContentType <= utils.VIDEO {
					// 保存消息只会在存在socket的一个端上进行保存，防止分布式部署后，消息重复问题
					_, exits := s.Clients[msg.From]
					if exits {
						saveMessage(msg)
					}

					if msg.MessageType == utils.MESSAGE_TYPE_USER {
						client, ok := s.Clients[msg.To]
						if ok {
							msgByte, err := json.Marshal(msg)
							if err == nil {
								client.Send <- msgByte
							}
						}
					} else if msg.MessageType == utils.MESSAGE_TYPE_GROUP {
						sendGroupMessage(msg, s)
					}
				} else {
					// 语音电话，视频电话等，仅支持单人聊天，不支持群聊
					// 不保存文件，直接进行转发
					client, ok := s.Clients[msg.To]
					if ok {
						client.Send <- message
					}
				}

			} else {
				// 无对应接受人员进行广播
				for id, conn := range s.Clients {
					zap.L().Info("allUser", zap.Any("allUser", id))

					select {
					case conn.Send <- message:
					default:
						close(conn.Send)
						delete(s.Clients, conn.Name)
					}
				}
			}
		}
	}
}

// 发送给群组消息,需要查询该群所有人员依次发送
func sendGroupMessage(msg *utils.Message, s *Server) {
	// 发送给群组的消息，查找该群所有的用户进行发送
	users := service.GroupService.GetUserIdByGroupUuid(msg.To)
	for _, user := range users {
		if user.Uuid == msg.From {
			continue
		}

		client, ok := s.Clients[user.Uuid]
		if !ok {
			continue
		}

		fromUserDetails := service.UserService.GetUserDetails(msg.From)
		// 由于发送群聊时，from是个人，to是群聊uuid。所以在返回消息时，将form修改为群聊uuid，和单聊进行统一
		msgSend := utils.Message{
			Avatar:       fromUserDetails.Avatar,
			FromUsername: msg.FromUsername,
			From:         msg.To,
			To:           msg.From,
			Content:      msg.Content,
			ContentType:  msg.ContentType,
			Type:         msg.Type,
			MessageType:  msg.MessageType,
			Url:          msg.Url,
		}

		msgByte, err := json.Marshal(&msgSend)
		if err == nil {
			client.Send <- msgByte
		}
	}
}

// 保存消息，如果是文本消息直接保存，如果是文件，语音等消息，保存文件后，保存对应的文件路径
func saveMessage(message *utils.Message) {
	// 如果上传的是base64字符串文件，解析文件保存
	if message.ContentType == 2 {
		url := uuid.New().String() + ".txt"

		content := message.Content

		dataBuffer, dataErr := base64.StdEncoding.DecodeString(content)
		if dataErr != nil {
			zap.L().Error("transfer base64 to file error", zap.String("transfer base64 to file error", dataErr.Error()))
			return
		}
		err := os.WriteFile(settings.Conf.FilePath+url, dataBuffer, 0666)
		if err != nil {
			zap.L().Error("write file error", zap.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.Content = ""
	} else if message.ContentType == 3 {
		// 普通的文件二进制上传
		fileSuffix := utils.GetFileType(message.File)
		nullStr := ""
		if nullStr == fileSuffix {
			fileSuffix = strings.ToLower(message.FileSuffix)
		}
		url := uuid.New().String() + "." + fileSuffix
		err := os.WriteFile(settings.Conf.FilePath+url, message.File, 0666)
		if err != nil {
			zap.L().Error("write file error", zap.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.File = nil
	}

	service.MessageService.SaveMessage(*message)
}

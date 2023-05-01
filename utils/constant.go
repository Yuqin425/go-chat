package utils

const (
	HEART_BEAT = "heartbeat"
	PONG       = "pong"

	// 消息类型，单聊或者群聊
	MESSAGE_TYPE_USER  = 1
	MESSAGE_TYPE_GROUP = 2

	// 消息内容类型
	TEXT  = 1
	FILE  = 2
	IMAGE = 3
	AUDIO = 4
	VIDEO = 5

	// 消息队列类型
	GO_CHANNEL = "gochannel"
	NSQ        = "nsq"
)

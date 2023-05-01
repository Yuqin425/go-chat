package utils

type MessageRequest struct {
	MessageType    int32  `json:"messageType"`
	Uuid           string `json:"uuid"`
	FriendUsername string `json:"friendUsername"`
}

type Message struct {
	Avatar       string `json:"avatar,omitempty" form:"avatar"`             //头像
	FromUsername string `json:"fromUsername,omitempty" form:"fromUsername"` // 发送消息用户的用户名
	From         string `json:"from,omitempty" form:"from"`                 // 发送消息用户uuid
	To           string `json:"to,omitempty" form:"to"`                     // 发送给对端用户的uuid
	Content      string `json:"content,omitempty" form:"content"`           // 文本消息内容
	ContentType  int32  `json:"contentType,omitempty" form:"contentType"`   // 消息内容类型：1.文字 2.普通文件 3.图片 4.音频 5.视频 6.语音聊天 7.视频聊天
	Type         string `json:"type,omitempty" form:"type"`                 // 消息传输类型：如果是心跳消息，该内容为heatbeat,在线视频或者音频为webrtc
	MessageType  int32  `json:"messageType,omitempty" form:"messageType"`   // 消息类型，1.单聊 2.群聊
	Url          string `json:"url,omitempty" form:"url"`                   // 图片，视频，语音的路径
	FileSuffix   string `json:"fileSuffix,omitempty" form:"fileSuffix"`     // 文件后缀，如果通过二进制头不能解析文件后缀，使用该后缀
	File         []byte `json:"file,omitempty" form:"file"`                 // 如果是图片，文件，视频等的二进制
}

type FriendRequest struct {
	Uuid           string `json:"uuid"`
	FriendUsername string `json:"friend_username"`
}

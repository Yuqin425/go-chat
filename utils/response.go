package utils

import (
	"IM_QQ/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

type GroupResponse struct {
	Uuid      string    `json:"uuid"`
	GroupId   int32     `json:"groupId"`
	CreatedAt time.Time `json:"createAt"`
	Name      string    `json:"name"`
	Notice    string    `json:"notice"`
}

type SearchResponse struct {
	User  models.User  `json:"user"`
	Group models.Group `json:"group"`
}

type MessageResponse struct {
	ID           int32     `json:"id" gorm:"primarykey"`
	FromUserId   int32     `json:"fromUserId" gorm:"index"`
	ToUserId     int32     `json:"toUserId" gorm:"index"`
	Content      string    `json:"content" gorm:"type:varchar(2500)"`
	ContentType  int16     `json:"contentType" gorm:"comment:'消息内容类型：1文字，2语音，3视频'"`
	CreatedAt    time.Time `json:"createAt"`
	FromUsername string    `json:"fromUsername"`
	ToUsername   string    `json:"toUsername"`
	Avatar       string    `json:"avatar"`
	Url          string    `json:"url"`
}

func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusInternalServerError, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusInternalServerError, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}

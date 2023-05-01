package utils

import "errors"

var (
	ErrorUserExit          = errors.New("用户已存在")
	ErrorFriendUserNotExit = errors.New("不存在此好友账号")
	ErrorFriendExit        = errors.New("已经是好友")
	ErrorPwdDifferent      = errors.New("两次密码不一致")
	ErrorUserNotExit       = errors.New("用户不存在")
	ErrorPasswordWrong     = errors.New("密码错误")
	ErrorsFriendNotExit    = errors.New("你和他还不是好友")
	ErrorInvalidID         = errors.New("无效的ID")
	ErrorInGroup           = errors.New("已加入群组")
	ErrorGroupNotExit      = errors.New("群组不存在")
	ErrorQueryFailed       = errors.New("查询数据失败")
	ErrorInsertFailed      = errors.New("插入数据失败")
)

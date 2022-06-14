package message

import (
	"wechat_server/api/application"
	"wechat_server/api/user"
)

const (
	// TypChatMsg 普通消息
	TypChatMsg = 1
	// TypApplicationMsg 申请加入room消息
	TypApplicationMsg = 2
	// TypApplicationResultMsg 处理申请结果
	TypApplicationResultMsg = 3
	// TypNewPeerAddedMsg 房间内有新人加入
	TypNewPeerAddedMsg = 4
)

type Message struct {
	// 消息类型
	Typ int `json:"typ,omitempty"`
	// 消息内容
	Content interface{} `json:"content"`
	// 发送者个人信息
	SendUser *user.User
}

func MakeChatMsg(content string,sendUser *user.User) Message {
	return Message{
		Typ:     TypChatMsg,
		Content: content,
		SendUser: sendUser,
	}
}

func MakeApplicationMsg(content application.Application,sendUser *user.User) Message {
	return Message{
		Typ:     TypApplicationMsg,
		Content: content,
		SendUser: sendUser,
	}
}

func MakeApplicationResultMsg(content application.Result,sendUser *user.User) Message {
	return Message{
		Typ:     TypApplicationResultMsg,
		Content: content,
		SendUser: sendUser,
	}
}

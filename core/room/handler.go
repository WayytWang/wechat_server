package room

import (
	"wechat_server/interactive/imodel"
)

type Handler interface {
	Handle(msg imodel.Message) error
}

func CreateHandler(msgType int) Handler {
	switch msgType {
	case imodel.TypChatMsg:
		return &ChatMsgHandler{}
	case imodel.TypApplicationMsg:
		return &ApplicationMsgHandler{}
	case imodel.TypApplicationResultMsg:
		return &ApplicationResultMsgHandler{}
	case imodel.TypNewPeerAddedMsg:
		return &ApplicationNewPeerAddedMsg{}
	}
	return nil
}

// ChatMsgHandler 聊天消息处理
type ChatMsgHandler struct {
}

func (h *ChatMsgHandler) Handle(msg imodel.Message) error {
	return nil
}

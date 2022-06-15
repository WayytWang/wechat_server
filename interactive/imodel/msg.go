package imodel

// 端对端之间交互消息体
// 消息包括：消息类型
//         消息内容（内容具体类型由消息类型决定）
//         发送者个人信息

const (
	// TypChatMsg 聊天消息
	TypChatMsg = 1
	// TypApplicationMsg 申请加入room消息
	TypApplicationMsg = 2
	// TypApplicationResultMsg 处理申请结果消息
	TypApplicationResultMsg = 3
	// TypNewPeerAddedMsg 新人加入消息
	TypNewPeerAddedMsg = 4
)

type Message struct {
	// 消息类型
	Typ int `json:"typ"`
	// 消息内容
	Content interface{} `json:"content"`
	// 发送者个人信息
	SendUser User `json:"send_user"`
}

func CreateChatMsg(content string) Message {
	return Message{
		Typ:     TypChatMsg,
		Content: content,
	}
}

func CreateApplicationMsg(content Application) Message {
	return Message{
		Typ:     TypApplicationMsg,
		Content: content,
	}
}

func CreateApplicationResultMsg(content Result) Message {
	return Message{
		Typ:     TypApplicationResultMsg,
		Content: content,
	}
}

func CreateNewPeerAddedMsg(content NewPeerAdded) Message{
	return Message{
		Typ:     TypNewPeerAddedMsg,
		Content: content,
	}
}
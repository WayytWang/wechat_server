package message

type Message struct {
	// 消息类型
	Typ int `json:"typ,omitempty"`
	// 消息内容
	Content interface{} `json:"content"`
}

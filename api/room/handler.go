package room

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"wechat_server/api/message"
)

// HandleNotify 循环监听tcp请求
func HandleNotify(ip string, port string) {
	address := fmt.Sprintf("%s:%s",ip,port)
	lister, err := net.Listen("tcp", address)
	if err != nil {
		panic(any(err))
	}
	defer lister.Close()

	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("连接Accept() 失败,err: ", err)
			continue
		}
		// 解析消息来源
		remoteAddr := conn.RemoteAddr().String()
		remoteSli := strings.Split(remoteAddr,":")
		remoteIp := remoteSli[0]
		// todo:remotePort为什么不对？
		remotePort := remoteSli[1]
		fmt.Println("remoteIp:"+remoteIp)
		fmt.Println("remotePort:"+remotePort)

		// 解析消息内容
		// todo:限制每条消息长度
		buf := make([]byte, 20000)
		length,err := conn.Read(buf)
		if err != nil {
			fmt.Printf("HandleNotify conn.Read(buf) 失败,err: %+v \n", err)
			continue
		}
		var msg message.Message
		err = json.Unmarshal(buf[:length],&msg)
		if err != nil {
			fmt.Printf("HandleNotify json.Unmarshal(buf,msg) 失败,err: %+v \n", err)
			continue
		}
		fmt.Printf("接收到内容:%+v", msg)

		// 校验
		if remoteIp != msg.SendUser.Ip {
			fmt.Println("HandleNotify 消息发送方伪造信息")
			continue
		}
		handler := MakeHandler(msg.Typ)
		err = handler.Handle(msg)
		if err != nil {
			fmt.Printf("HandleNotify Handle 失败,err: %+v \n", err)
			continue
		}
	}
}


type Handler interface {
	Handle(msg message.Message) error
}

func MakeHandler(msgType int) Handler {
	switch msgType {
	case message.TypChatMsg:
		return &ChatMsgHandler{}
	case message.TypApplicationMsg:
		return &ApplicationMsgHandler{}
	case message.TypApplicationResultMsg:
		return &ApplicationResultMsgHandler{}
	}
	return nil
}

// ChatMsgHandler 聊天消息处理
type ChatMsgHandler struct {
}

func (h *ChatMsgHandler) Handle(msg message.Message) error {
	return nil
}



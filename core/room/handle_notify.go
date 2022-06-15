package room

import (
	"fmt"
	"net"
	"wechat_server/core/model"
	"wechat_server/tcp_conn"
)

// HandleNotify 循环监听tcp请求 todo:错误处理
func HandleNotify(r *model.Room) {
	address := fmt.Sprintf("%s:%s", r.Listener.Ip, r.Listener.Port)
	lister, err := net.Listen("tcp", address)
	if err != nil {
		panic(any(err))
	}
	r.Lister = lister
	defer lister.Close()

	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("连接Accept() 失败,err: ", err)
			continue
		}
		tc := tcp_conn.CreateTcpConn(conn)
		msg, err := tc.ParseMsg()
		if err != nil {
			fmt.Printf("[Room] [HandleNotify] tc.ParseMsg() error: %+v \n", err)
			continue
		}
		h := CreateHandler(msg.Typ)
		err = h.Handle(msg)
		if err != nil {
			fmt.Printf("[Room] [HandleNotify]  handler.Handle(msg) error: %+v \n", err)
			continue
		}
	}
}

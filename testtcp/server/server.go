package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		// 等待客户端通过conn发送消息
		// 如果没有消息就会阻塞与此
		fmt.Printf("服务器等待客户端%s 发送信息 \n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端已退出,err:", err)
			return
		} else {
			fmt.Printf("收到了客户端 %s 数据:%s ", conn.RemoteAddr().String(), string(buf[:n]))
		}
	}
}

func main() {
	fmt.Println("服务器开始监听...")

	lister, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败...err:", err)
		return
	}
	defer lister.Close()

	for {
		fmt.Println("等待客户端连接")
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("连接Accept() 失败,err: ", err)
		} else {
			fmt.Printf("Accept() suc conn=%v,客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}

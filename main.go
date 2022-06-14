package main

import (
	"os"
	"os/signal"
	"syscall"
	"wechat_server/command"
)

func main() {
	go command.Execute()
	// 阻塞住程序
	// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前 Goroutine 等待信号
	<-quit
}

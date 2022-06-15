package convert

import (
	"wechat_server/core/application"
	"wechat_server/core/user"
	"wechat_server/interactive/imodel"
)

// 交互结构体 to 运行时结构体

func FromIUser(u imodel.User) user.User {
	return user.User{
		Name: u.Name,
		Ip:   u.Ip,
		Port: u.Port,
	}
}

func FromIApplication(a imodel.Application) application.Application {
	return application.Application{
		ApplyRoomID:  a.ApplyRoomID,
		ApplyContent: a.ApplyContent,
	}
}

func FromIApplicationResult(r imodel.Result) application.Result {
	return application.Result{
		IsOk: r.IsOk,
	}
}
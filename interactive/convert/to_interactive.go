package convert

import (
	"wechat_server/core/user"
	"wechat_server/interactive/imodel"
)

// 运行时结构体 to 交互结构体

func FromUser(u *user.User) imodel.User {
	return imodel.User{
		Name: u.Name,
		Ip:   u.Ip,
		Port: u.Port,
	}
}

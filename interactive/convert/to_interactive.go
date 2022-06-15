package convert

import (
	"wechat_server/core/model"
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

func FromUsers(us []*user.User) []imodel.User {
	iUs := make([]imodel.User, 0)
	for i := range us {
		u := FromUser(us[i])
		iUs = append(iUs, u)
	}
	return iUs
}

func FromRoom(r *model.Room) imodel.Room {
	return imodel.Room{
		ID:      r.ID,
		Name:    r.Name,
		Creator: FromUser(r.Creator),
		Peers:   FromUsers(r.Peers),
	}
}

package convert

import (
	"wechat_server/core/room"
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

func FromIUsers(iUs []imodel.User) []*user.User {
	us := make([]*user.User,0)
	for i := range iUs {
		u := FromIUser(iUs[i])
		us = append(us, &u)
	}
	return us
}

func FromIApplication(a imodel.Application) room.Application {
	return room.Application{
		ApplyRoomID:  a.ApplyRoomID,
		ApplyContent: a.ApplyContent,
	}
}

func FromIRoom(r imodel.Room) room.Room {
	c := FromIUser(r.Creator)
	return room.Room{
		ID:       r.ID,
		Name:     r.Name,
		Creator:  &c,
		Peers: FromIUsers(r.Peers),
	}
}

func FromIApplicationResult(r imodel.Result) room.Result {
	rm := FromIRoom(r.Room)
	return room.Result{
		Room: &rm,
		IsOk: r.IsOk,
	}
}
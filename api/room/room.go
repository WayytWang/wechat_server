package room

import (
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"wechat_server/api/user"
	"wechat_server/cache"
)

type Room struct {
	ID string
	Name string

	Creator *user.User
	// 房间成员
	peers []*user.User
}

func CreatRoom(roomName string,userName string) *Room {
	// 获取房间创建者信息(公网ip)
	myInfo := cache.GetMyInfo()
	if myInfo.IpAddress == "" {
		panic(any(errors.New("count not access ip address")))
	}
	u := uuid.NewV4()
	id := u.String()
	r := &Room{
		ID:      id,
		Creator: myInfo,
		peers:   make([]*user.User,0),
	}
	cache.GetRoomMap().AddRoom(id,r)

	// todo: 启动"room"协程
	return r
}

func (r *Room) AddPeers(peer *user.User) {
	// todo: notify other peers, a new peer will be added this room
	r.peers = append(r.peers, peer)
}

func (r *Room) GetPeers() []*user.User {
	return r.peers
}


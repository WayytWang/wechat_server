package model

import (
	"net"
	"sync"
	"wechat_server/core/user"
)

type Room struct {
	ID   string
	Name string

	Creator *user.User
	Listener *user.User
	Lister   net.Listener

	mutex sync.RWMutex
	// 房间成员 不包括自己
	Peers []*user.User
}

// InitRoom 创建房间时调用
func InitRoom(id, name string, creator *user.User) *Room {
	return &Room{
		ID:       id,
		Name:     name,
		Creator:  creator,
		Listener: creator,
		mutex:    sync.RWMutex{},
		Peers:    make([]*user.User, 0),
	}
}

// InitEmptyRoom 被批准加入房间前内存中保存的房间信息
func InitEmptyRoom(id string,listener *user.User) *Room {
	return &Room{
		ID:       id,
		Listener: listener,
		mutex:    sync.RWMutex{},
		Peers:    make([]*user.User, 0),
	}
}

// AddPeers 增加节点信息
func (r *Room) AddPeers(peer *user.User) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.Peers = append(r.Peers, peer)
}

func (r *Room) GetPeers() []*user.User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.Peers
}

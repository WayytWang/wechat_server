package room

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
	lister  net.Listener

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

//// SendMessage 发送消息
//func (r *Room) SendMessage(sendUser *user.User, message message.Message) error {
//	conn, err := tcp_conn.GetConnByUser(&tcp_conn.TcpInfo{
//		Ip:   sendUser.Ip,
//		Port: sendUser.Port,
//	})
//	bytes, err := json.Marshal(message)
//	if err != nil {
//		err = errors.Wrap(err, "[Room.SendMessage] json.Marshal(message) error")
//		return err
//	}
//	_, err = conn.Write(bytes)
//	if err != nil {
//		err = errors.Wrap(err, "[Room.SendMessage] Write error")
//		return err
//	}
//	return nil
//}

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

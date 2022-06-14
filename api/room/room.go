package room

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net"
	"wechat_server/api/application"
	"wechat_server/api/message"
	"wechat_server/api/user"
	"wechat_server/cache"
)

type Room struct {
	ID   string
	Name string

	Creator *user.User
	// 房间成员
	peers []*user.User

	lister net.Listener
}

// JoinRoomApplication 申请加入房间
func JoinRoomApplication(shareStr string,joinMessage string) (string,error) {
	shareRoomStr := ShareRoomStr(shareStr)
	ip,port,roomID := shareRoomStr.Parse()
	// 获取与房主的tcp连接
	conn,err := cache.GetConnByUser(&cache.TcpInfo{
		Ip:   ip,
		Port: port,
	})
	if err != nil {
		return "",err
	}
	applyInfo := application.CreateApplication(user.GetMyInfo(),roomID,joinMessage)
	applyMsg := message.MakeApplicationMsg(applyInfo,user.GetMyInfo())
	applyBytes,err := json.Marshal(applyMsg)
	if err != nil {
		return "", errors.Wrap(err,"[JoinRoomApplication] json.Marshal(applyMsg) error")
	}
	// 构造申请信息
	_,err = conn.Write(applyBytes)
	if err != nil {
		return "", errors.Wrap(err,"[JoinRoomApplication] conn.Write(applyBytes) error")
	}
	// 阻塞等待房主审批
	go HandleNotify(user.GetMyInfo().Ip,user.GetMyInfo().Port)
	//return roomID,nil
	return "",nil
}

// SendMessage 发送消息
func (r *Room) SendMessage(sendUser *user.User, message message.Message) error {
	conn, err := cache.GetConnByUser(&cache.TcpInfo{
		Ip:   sendUser.Ip,
		Port: sendUser.Port,
	})
	bytes, err := json.Marshal(message)
	if err != nil {
		err = errors.Wrap(err, "[Room.SendMessage] json.Marshal(message) error")
		return err
	}
	_, err = conn.Write(bytes)
	if err != nil {
		err = errors.Wrap(err, "[Room.SendMessage] Write error")
		return err
	}
	return nil
}

func (r *Room) AddPeers(peer *user.User) {
	// todo: notify other peers, a new peer will be added this room
	r.peers = append(r.peers, peer)
}

func (r *Room) GetPeers() []*user.User {
	return r.peers
}

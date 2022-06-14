package room

import (
	"fmt"
	"github.com/pkg/errors"
	"wechat_server/api/user"
)

func CreatRoom(roomName string) *Room {
	// 获取房间创建者信息(公网ip)
	myInfo := user.GetMyInfo()
	if myInfo.Ip == "" {
		panic(any(errors.New("count not gain ip address")))
	}
	fmt.Println(myInfo)
	// todo：测试阶段随机id
	//u := uuid.NewV4()
	//id := u.String()
	id := "test"
	// 启动tcp server 监听"room"信息
	go HandleNotify(myInfo.Ip,myInfo.Port)
	r := &Room{
		ID:      id,
		Creator: myInfo,
		Name:    roomName,
		peers:   make([]*user.User, 0),
	}
	GetRoomMap().AddRoom(id, r)
	shareRoomStr := CreateShareRoomStr(myInfo.Ip,myInfo.Port,id)
	fmt.Printf("您的房间[%s]创建成功,请复制[%s]分享给您的好友\n", roomName, shareRoomStr)
	return r
}

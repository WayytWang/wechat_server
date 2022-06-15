package core

import (
	"fmt"
	"github.com/pkg/errors"
	"wechat_server/core/model"
	"wechat_server/core/room"
	"wechat_server/core/user"
	"wechat_server/utils"
)

func CreatRoom(roomName string) *model.Room {
	// 获取自己的信息
	myInfo := user.GetMyInfo()
	if myInfo.Ip == "" {
		panic(any(errors.New("count not gain ip address")))
	}
	fmt.Println(myInfo)
	// todo：测试阶段随机id
	//u := uuid.NewV4()
	//id := u.String()
	id := "test"
	r := model.InitRoom(id, roomName, myInfo)
	// 启动tcp server 监听"room"信息
	go room.HandleNotify(r)
	model.GetRoomMap().AddRoom(id, r)
	shareRoomStr := CreateShareRoomStr(myInfo.Ip, myInfo.Port, id)
	utils.TipsPrint(fmt.Sprintf("您的房间[%s]创建成功,请复制[%s]分享给您的好友", roomName, shareRoomStr))
	return r
}

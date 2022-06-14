package application

import (
	"wechat_server/api/user"
)

// Application 申请加入房间
type Application struct {
	// 申请人信息
	ApplyUser *user.User
	// 申请加入的房间ID
	ApplyRoomID string


	ApplyMessage string
}

func CreateApplication(applyUser *user.User,roomID string, applyMsg string) Application {
	return Application{
		ApplyUser:    applyUser,
		ApplyRoomID:  roomID,
		ApplyMessage: applyMsg,
	}
}

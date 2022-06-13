package application

import "wechat_server/user"

// Application 申请加入房间
type Application struct {
	ApplyUser *user.User
	ApplyRoom string

	ApplyMessage string
}

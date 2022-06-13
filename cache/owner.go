package cache

import (
	uuid "github.com/satori/go.uuid"
	"sync"
	"wechat_server/user"
	"wechat_server/utils"
)

var (
	myInfo *user.User
	myInfoOnce = &sync.Once{}
)

func SetMyInfo(name string, port int) {
	// 用户id
	u := uuid.NewV4()
	id := u.String()

	// 用户ip
	ip,err := utils.ExternalIP()
	if err != nil {
		panic(any(err))
	}

	myInfoOnce.Do(func() {
		myInfo = &user.User{
			ID:        id,
			Name: name,
			IpAddress: string(ip),
			Port:      port,
		}
	})
}

func GetMyInfo() *user.User {
	return myInfo
}

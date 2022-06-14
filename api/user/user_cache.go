package user

import (
	"sync"
	"wechat_server/utils"
)

var (
	myInfo     *User
	myInfoOnce = &sync.Once{}
)

func SetMyInfo(name string, port string) {
	myInfoOnce.Do(func() {
		// 用户ip
		ip, err := utils.ExternalIP()
		if err != nil {
			panic(any(err))
		}
		myInfo = &User{
			Name: name,
			Ip:   ip,
			Port: port,
		}
	})
}

func GetMyInfo() *User {
	return myInfo
}

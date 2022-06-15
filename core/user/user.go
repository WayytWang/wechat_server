package user

import "fmt"

type User struct {
	Name string

	Ip   string
	Port string
}

func (u User) String() string {
	info := fmt.Sprintf("您的个人信息: \n昵称:%s \nip地址:%s \n网络端口:%s \n", u.Name, u.Ip, u.Port)
	return info
}

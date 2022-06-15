package imodel

type Room struct {
	ID   string
	Name string

	// 创建者
	Creator User
	// 房间成员
	peers []User
}

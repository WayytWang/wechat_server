package core

import (
	"fmt"
	"strings"
)

// ShareRoomStr 分享字符串 todo:考虑分享加密后的字符串
type ShareRoomStr string

func CreateShareRoomStr(ip string, port string, roomId string) ShareRoomStr {
	return ShareRoomStr(fmt.Sprintf("%s#%s#%s", ip, port, roomId))
}

func (s ShareRoomStr) Parse() (ip string, port string, roomID string) {
	result := strings.Split(string(s), "#")
	return result[0], result[1], result[2]
}

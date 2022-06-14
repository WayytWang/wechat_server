package room

import (
	"fmt"
	"strings"
)

type ShareRoomStr string

func CreateShareRoomStr(ip string,port string,roomId string) ShareRoomStr {
	return ShareRoomStr(fmt.Sprintf("%s#%s#%s",ip,port,roomId))
}

func (s ShareRoomStr) Parse() (ip string,port string,roomID string) {
	result := strings.Split(string(s),"#")
	return result[0],result[1],result[2]
}

package cache

import (
	"sync"
	"wechat_server/api/room"
)

// typeRoomMap 房间信息保存在内存中
type typeRoomMap struct {
	rm map[string]*room.Room
}

// AddRoom 添加房间信息至内存中
func (rm *typeRoomMap) AddRoom(id string, room *room.Room) {
	rm.rm[id] = room
}

var (
	roomOnce = &sync.Once{}
	roomMap *typeRoomMap
)

func GetRoomMap() *typeRoomMap {
	roomOnce.Do(func() {
		roomMap = &typeRoomMap{
			rm:make(map[string]*room.Room),
		}
	})
	return roomMap
}




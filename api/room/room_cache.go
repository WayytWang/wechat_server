package room

import (
	"sync"
)

// typeRoomMap 房间信息保存在内存中
type typeRoomMap struct {
	rm map[string]*Room
}

// AddRoom 添加房间信息至内存中
func (rm *typeRoomMap) AddRoom(id string, room *Room) {
	rm.rm[id] = room
}
// GetRoom 添加房间信息至内存中
func (rm *typeRoomMap) GetRoom(id string)(room *Room) {
	room = rm.rm[id]
	return
}


var (
	roomOnce = &sync.Once{}
	roomMap  *typeRoomMap
)

func GetRoomMap() *typeRoomMap {
	roomOnce.Do(func() {
		roomMap = &typeRoomMap{
			rm: make(map[string]*Room),
		}
	})
	return roomMap
}

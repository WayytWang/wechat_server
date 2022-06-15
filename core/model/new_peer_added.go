package model

import (
	"wechat_server/core/user"
)

type NewPeerAdded struct {
	RoomID string     `json:"room_id"`
	NewPeer *user.User `json:"new_peer"`
	Peers   []*user.User `json:"peers"`
}

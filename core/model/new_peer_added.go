package model

import "wechat_server/core/user"

type NewPeerAdded struct {
	Peers []*user.User `json:"peers"`
}

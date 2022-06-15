package imodel

type NewPeerAdded struct {
	RoomID string `json:"room_id"`
	NewPeer User `json:"new_peer"`
	Peers   []User   `json:"peers"`
}

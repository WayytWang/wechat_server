package room

type Result struct {
	Room *Room `json:"room"`
	IsOk bool `json:"is_ok"`
}

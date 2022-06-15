package room

import "wechat_server/interactive/imodel"

type ApplicationNewPeerAddedMsg struct {
}

func (h *ApplicationNewPeerAddedMsg) Handle(msg imodel.Message) error {
	return nil
}

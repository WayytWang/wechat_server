package room

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"wechat_server/core/model"
	"wechat_server/core/user"
	"wechat_server/interactive/convert"
	"wechat_server/interactive/imodel"
	"wechat_server/utils"
)

type NewPeerAddedHandler struct {
}

func (h *NewPeerAddedHandler) Handle(msg imodel.Message) error {
	npa,err := parseNewPeerAddedMsg(msg)
	if err != nil {
		return err
	}
	r := model.GetRoomMap().GetRoom(npa.RoomID)
	for _, peer := range npa.Peers {
		if peer == nil {
			continue
		}
		if peer.Name == user.GetMyInfo().Name {
			continue
		}
		r.Peers = append(r.Peers, peer)
	}
	r.Peers = append(r.Peers,npa.NewPeer)
	peerStr := ""
	for i,p := range r.Peers {
		peerStr += p.Name
		if i < len(r.Peers)-1 {
			peerStr += ","
		}
	}
	model.GetRoomMap().AddRoom(npa.RoomID,r)
	printStr := fmt.Sprintf("[%s]加入的房间[%s],房间内现有成员[%s]",npa.NewPeer.Name,r.Name,peerStr)
	utils.TipsPrint(printStr)
	return nil
}

func parseNewPeerAddedMsg(msg imodel.Message) (npa model.NewPeerAdded, err error) {
	// 解析消息内容
	bytes, err := json.Marshal(msg.Content)
	if err != nil {
		err = errors.Wrap(err, "[NewPeerAddedHandler] [parseNewPeerAddedMsg] json.Marshal(msg.Content)")
		return npa, err
	}
	var iNpa imodel.NewPeerAdded
	err = json.Unmarshal(bytes, &iNpa)
	if err != nil {
		err = errors.Wrap(err, "[NewPeerAddedHandler] [parseNewPeerAddedMsg] json.Unmarshal(bytes,&apply)")
		return npa, err
	}
	npa = convert.FromINewPeerAdded(iNpa)
	return
}

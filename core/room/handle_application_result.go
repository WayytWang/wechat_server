package room

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"wechat_server/core/model"
	"wechat_server/interactive/convert"
	"wechat_server/interactive/imodel"
	"wechat_server/utils"
)

// ApplicationResultMsgHandler 申请加入room消息处理
type ApplicationResultMsgHandler struct {
}

func (h *ApplicationResultMsgHandler) Handle(msg imodel.Message) error {
	result, err := parseApplicationResultMsg(msg)
	if err != nil {
		return err
	}
	if !result.IsOk {
		utils.TipsPrint("您被拒绝了")
		os.Exit(1)
	}
	// 成功加入,更新内存
	r := result.Room
	cr := model.GetRoomMap().GetRoom(r.ID)
	cr.Creator = r.Creator
	cr.Peers = r.Peers
	cr.AddPeers(r.Creator)
	cr.Name = r.Name
	model.GetRoomMap().AddRoom(r.ID, cr)
	peerStr := ""
	for i, peer := range cr.GetPeers() {
		peerStr += peer.Name
		if i < len(cr.GetPeers())-1 {
			peerStr += ","
		}
	}
	printStr := fmt.Sprintf("您已成功加入房间[%s],房间内现有成员[%s]", cr.Name,peerStr)
	utils.TipsPrint(printStr)
	return nil
}

func parseApplicationResultMsg(msg imodel.Message) (result model.Result, err error) {
	// 解析消息内容
	bytes, err := json.Marshal(msg.Content)
	if err != nil {
		err = errors.Wrap(err, "[ApplicationResultMsgHandler] [parseApplicationResultMsg] json.Marshal(msg.Content)")
		return result, err
	}
	var iResult imodel.Result
	err = json.Unmarshal(bytes, &iResult)
	if err != nil {
		err = errors.Wrap(err, "[ApplicationResultMsgHandler] [parseApplicationResultMsg] json.Unmarshal(bytes,&apply)")
		return result, err
	}
	result = convert.FromIApplicationResult(iResult)
	return
}

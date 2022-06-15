package room

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"wechat_server/core/application"
	"wechat_server/interactive/convert"
	"wechat_server/interactive/imodel"
	"wechat_server/utils"
)

// ApplicationResultMsgHandler 申请加入room消息处理
type ApplicationResultMsgHandler struct {
}

func (h *ApplicationResultMsgHandler) Handle(msg imodel.Message) error {
	result,err := parseApplicationResultMsgHandler(msg)
	if err != nil {
		return err
	}
	if !result.IsOk {
		utils.TipsPrint("您被拒绝了")
		os.Exit(1)
	}
	// 成功加入,更新内存
	return nil
}

func parseApplicationResultMsgHandler(msg imodel.Message) (result application.Result,err error) {
	// 解析消息内容
	bytes, err := json.Marshal(msg.Content)
	if err != nil {
		err = errors.Wrap(err, "[ApplicationResultMsgHandler] [parseApplicationResultMsgHandler] json.Marshal(msg.Content)")
		return result,err
	}
	var iResult imodel.Result
	err = json.Unmarshal(bytes, &iResult)
	if err != nil {
		err = errors.Wrap(err, "[ApplicationResultMsgHandler] [parseApplicationResultMsgHandler] json.Unmarshal(bytes,&apply)")
		return result,err
	}
	result = convert.FromIApplicationResult(iResult)
	return
}

package room

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"wechat_server/api/application"
	"wechat_server/api/message"
)

// ApplicationResultMsgHandler 申请加入room消息处理
type ApplicationResultMsgHandler struct {
}

func (h *ApplicationResultMsgHandler) Handle(msg message.Message) error {
	bytes,err := json.Marshal(msg.Content)
	if err != nil {
		err = errors.Wrap(err,"[ApplicationResultMsgHandler] [Handle] json.Marshal(msg.Content)")
		return err
	}
	var result application.Result
	err = json.Unmarshal(bytes,&result)
	if err != nil {
		err = errors.Wrap(err,"[ApplicationResultMsgHandler] [Handle] json.Unmarshal(bytes,&result)")
		return nil
	}
	if !result.IsOk {
		fmt.Println("您被拒绝了")
		return nil
	}
	return nil
}

package room

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"wechat_server/api/application"
	"wechat_server/api/message"
	"wechat_server/api/user"
	"wechat_server/cache"
)

// ApplicationMsgHandler 申请加入room消息处理
type ApplicationMsgHandler struct {
}

func (h *ApplicationMsgHandler) Handle(msg message.Message) error {
	bytes,err := json.Marshal(msg.Content)
	if err != nil {
		err = errors.Wrap(err,"[ApplicationMsgHandler] [Handle] json.Marshal(msg.Content)")
		return err
	}
	var apply application.Application
	err = json.Unmarshal(bytes,&apply)
	if err != nil {
		err = errors.Wrap(err,"[ApplicationMsgHandler] [Handle] json.Unmarshal(bytes,&apply)")
		return nil
	}
	// 获取房间信息
	roomInfo := GetRoomMap().GetRoom(apply.ApplyRoomID)
	if roomInfo == nil {
		fmt.Printf("申请加入房间的房间号[%s]不存在 \n",apply.ApplyRoomID)
		return nil
	}
	displayStr := fmt.Sprintf("%s[%s] 申请加入房间[%s],申请信息[%s],是否同意？\n yes:同意 no:让他滚",
		msg.SendUser.Name,msg.SendUser.Ip,roomInfo.Name,apply.ApplyMessage)
	fmt.Println(displayStr)
	// 等待终端输入指令
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return nil
	}
	line = strings.Trim(line, " \r\n")
	fmt.Println("审批意见:",line)
	// 不同意
	if line == "no" {
		// 通知申请人
		conn,err := cache.GetConnByUser(&cache.TcpInfo{
			Ip:   msg.SendUser.Ip,
			Port: msg.SendUser.Port,
		})
		if err != nil {
			return err
		}
		result := application.Result{
			IsOk: false,
		}
		msg := message.MakeApplicationResultMsg(result,user.GetMyInfo())
		bytes,err := json.Marshal(msg)
		if err != nil {
			return err
		}
		_,err = conn.Write(bytes)
		if err != nil {
			return err
		}
		return nil
	}
	// 同意
	if line == "yes" {
		// todo:通知所有人，房间有新节点加入
	}
	return nil
}

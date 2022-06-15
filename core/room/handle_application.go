package room

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"wechat_server/core/model"
	"wechat_server/interactive/convert"
	"wechat_server/interactive/imodel"
	"wechat_server/tcp_conn"
	"wechat_server/utils"
)

// ApplicationMsgHandler 申请加入room消息处理
type ApplicationMsgHandler struct {
}

func (h *ApplicationMsgHandler) Handle(msg imodel.Message) error {
	// 解析消息内容
	apply,err := parseApplicationMsg(msg)
	if err != nil {
		return err
	}
	// 用户消息
	callerUser := msg.SendUser
	// 获取房间信息
	roomInfo := model.GetRoomMap().GetRoom(apply.ApplyRoomID)
	if roomInfo == nil {
		utils.TipsPrint(fmt.Sprintf("申请加入房间的房间号[%s]不存在", apply.ApplyRoomID))
		return nil
	}
	displayStr := fmt.Sprintf("%s[%s] 申请加入房间[%s],申请信息[%s],是否同意？\n yes:同意 no:让他滚",
		msg.SendUser.Name, msg.SendUser.Ip, roomInfo.Name, apply.ApplyContent)
	utils.TipsPrint(displayStr)
	// 等待终端输入指令
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return nil
	}
	line = strings.Trim(line, " \r\n")
	// 不同意
	if line == "no" {
		result := imodel.Result{
			IsOk: false,
		}
		// 通知申请人
		backMsg := imodel.CreateApplicationResultMsg(result)
		return tcp_conn.TcpSendMsg(callerUser.Ip, callerUser.Port, backMsg)
	}
	// 同意
	if line == "yes" {
		// 通知申请者成功加入
		result := imodel.Result{
			Room: convert.FromRoom(roomInfo),
			IsOk: true,
		}
		backMsg := imodel.CreateApplicationResultMsg(result)
		err = tcp_conn.TcpSendMsg(callerUser.Ip, callerUser.Port, backMsg)
		if err != nil {
			return err
		}
		// 将申请者信息加入room Peers
		// roomInfo.AddPeers(callerUser)
		// 通知所有人，房间有新节点加入

	}
	return nil
}

func parseApplicationMsg(msg imodel.Message) (apply model.Application,err error) {
	// 解析消息内容
	bytes, err := json.Marshal(msg.Content)
	if err != nil {
		err = errors.Wrap(err, "[ApplicationMsgHandler] [parseApplicationMsg] json.Marshal(msg.Content)")
		return apply,err
	}
	var iApply imodel.Application
	err = json.Unmarshal(bytes, &iApply)
	if err != nil {
		err = errors.Wrap(err, "[ApplicationMsgHandler] [parseApplicationMsg] json.Unmarshal(bytes,&apply)")
		return apply,err
	}
	apply = convert.FromIApplication(iApply)
	return
}

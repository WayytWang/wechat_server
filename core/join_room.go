package core

import (
	"wechat_server/core/model"
	"wechat_server/core/room"
	"wechat_server/core/user"
	"wechat_server/interactive/imodel"
	"wechat_server/tcp_conn"
	"wechat_server/utils"
)

// JoinRoomApplication 申请加入房间
func JoinRoomApplication(shareStr string, joinMessage string) (string, error) {
	shareRoomStr := ShareRoomStr(shareStr)
	ip, port, roomID := shareRoomStr.Parse()
	// 获取与房主的tcp连接
	myInfo := user.GetMyInfo()
	r := model.InitEmptyRoom(roomID,myInfo)
	// 构建消息内容
	application := imodel.CreateApplication(roomID,joinMessage)
	applyMsg := imodel.CreateApplicationMsg(application)
	// 给房主发送加入申请消息
	err := tcp_conn.TcpSendMsg(ip,port,applyMsg)
	if err != nil {
		return "", err
	}
	utils.TipsPrint("等待房主的审批结果")
	go room.HandleNotify(r)
	model.GetRoomMap().AddRoom(r.ID,r)
	return "", nil
}

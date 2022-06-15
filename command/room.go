package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"wechat_server/core"
	"wechat_server/core/user"
)

// 创建房间参数
var (
	roomName string
	yourName string
	yourPort string
)

// 加入房间参数
var (
	share      string
	joinerName string
	joinMsg    string
	joinerPort string
)

func InitRoomCommand() *cobra.Command {
	RoomCommand.AddCommand(RoomCreateCommand)
	RoomCommand.AddCommand(RoomJoinCommand)
	return RoomCommand
}

var RoomCommand = &cobra.Command{
	Use:     "room",
	Short:   "房间相关操作:创建房间、加入房间",
	Example: "",
	RunE: func(c *cobra.Command, args []string) error {
		return nil
	},
}

var RoomCreateCommand = &cobra.Command{
	Use:     "create",
	Short:   "输入房间名、您的昵称、网络端口来创建房间",
	Aliases: []string{"c"},
	Example: "wechat room create 相亲相爱一家人 知足常乐 8888",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("创建房间参数必须包括房间名、房主名、网络端口")
		}
		roomName = args[0]
		yourName = args[1]
		yourPort = args[2]
		return nil
	},
	RunE: func(c *cobra.Command, args []string) error {
		user.SetMyInfo(yourName, yourPort)
		core.CreatRoom(roomName)
		return nil
	},
}

var RoomJoinCommand = &cobra.Command{
	Use:     "join",
	Short:   "输入房主分享信息以及个人昵称以及验证信息加入聊天房间",
	Aliases: []string{"j"},
	Example: "wechat room join shareString 陈冠希 我陈冠希想加入你们",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 4 {
			return errors.New("加入房间参数必须包括房主分享信息、昵称、验证信息、网络端口")
		}
		share = args[0]
		joinerName = args[1]
		joinMsg = args[2]
		joinerPort = args[3]
		return nil
	},
	RunE: func(c *cobra.Command, args []string) error {
		user.SetMyInfo(joinerName, joinerPort)
		roomName, err := core.JoinRoomApplication(share, joinMsg)
		if err != nil {
			return err
		}
		if roomName != "" {
			fmt.Printf("您已成功加入房间[%s],可以开始聊天 \n", roomName)
		}
		return nil
	},
}

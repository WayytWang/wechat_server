package command

import (
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(InitRoomCommand())
}

var rootCmd = &cobra.Command{
	Use:   "wechat",
	Short: "端对端聊天,无渠道保存任何个人信息和聊天内容,想喷就喷,无负担聊天",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package command

import (
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(InitExecCommand())
}

var rootCmd = &cobra.Command{
	Use:   "wechat [COMMANDS]",
	Short: "文档库及目录创建工具",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
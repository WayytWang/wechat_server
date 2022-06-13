package command

import (
	"github.com/spf13/cobra"
)


var execFileName string

func InitExecCommand() *cobra.Command {
	ExecCommand.PersistentFlags().StringVar(&execFileName, "file", "gen.xlsx", "需要处理的文件名")
	return ExecCommand
}

var ExecCommand = &cobra.Command{
	Use:     "exec [flags]",
	Short:   "根据xlsx文件生成文档库及目录",
	Aliases: []string{"e"},
	Example: "",
	RunE: func(c *cobra.Command, args []string) error {
		return nil
	},
}
package cmd

import (
	"github.com/jeanhua/ZanaoCLI/internal/output"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "用户相关操作",
}

var userMessagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "获取消息列表",
	RunE: func(cmd *cobra.Command, args []string) error {
		msgs, err := getClient().GetMessage()
		if err != nil {
			return err
		}
		output.PrintItems(*msgs, jsonOutput)
		return nil
	},
}

var userInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "获取当前用户信息",
	RunE: func(cmd *cobra.Command, args []string) error {
		info, err := getClient().GetUserInfo()
		if err != nil {
			return err
		}
		output.PrintItem(*info, jsonOutput)
		return nil
	},
}

func init() {
	userCmd.AddCommand(userMessagesCmd)
	userCmd.AddCommand(userInfoCmd)
}

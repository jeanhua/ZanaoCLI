package cmd

import (
	"fmt"
	"os"

	"github.com/jeanhua/ZanaoCLI/zanao"
	"github.com/spf13/cobra"
)

var jsonOutput bool

var rootCmd = &cobra.Command{
	Use:   "zanao",
	Short: "赞哦校园集市 CLI",
	Long:  "赞哦校园集市命令行工具，支持浏览、发布、评论等操作。",
	Example: `  zanao posts hot
  zanao posts search --keyword '二手书'
  zanao user info
  zanao posts hot --json`,
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("ZANAO_TOKEN") == "" {
			return fmt.Errorf("ZANAO_TOKEN 环境变量未设置\n从微信小程序抓包的 X-Sc-Od 请求头中获取 token")
		}
		if os.Getenv("ZANAO_SCHOOL_ALIAS") == "" {
			return fmt.Errorf("ZANAO_SCHOOL_ALIAS 环境变量未设置\n例如：export ZANAO_SCHOOL_ALIAS=scu")
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "以 JSON 格式输出")
	rootCmd.AddCommand(postsCmd)
	rootCmd.AddCommand(commentCmd)
	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(categoriesCmd)
}

func getClient() *zanao.ZanaoClient {
	return zanao.NewZanaoClient(
		os.Getenv("ZANAO_TOKEN"),
		os.Getenv("ZANAO_SCHOOL_ALIAS"),
	)
}

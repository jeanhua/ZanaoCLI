package cmd

import (
	"github.com/jeanhua/ZanaoCLI/internal/output"
	"github.com/spf13/cobra"
)

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "获取帖子分类列表",
	RunE: func(cmd *cobra.Command, args []string) error {
		cats, err := getClient().GetCategory()
		if err != nil {
			return err
		}
		output.PrintItems(*cats, jsonOutput)
		return nil
	},
}

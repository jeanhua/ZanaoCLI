package cmd

import (
	"github.com/jeanhua/ZanaoCLI/internal/output"
	"github.com/spf13/cobra"
)

var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "评论相关操作",
}

var commentListCmd = &cobra.Command{
	Use:   "list",
	Short: "获取帖子的评论列表",
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, _ := cmd.Flags().GetString("thread")
		comments, err := getClient().GetComment(thread)
		if err != nil {
			return err
		}
		output.PrintItems(*comments, jsonOutput)
		return nil
	},
}

var commentPostCmd = &cobra.Command{
	Use:   "post",
	Short: "发表评论",
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, _ := cmd.Flags().GetString("thread")
		content, _ := cmd.Flags().GetString("content")
		reply, _ := cmd.Flags().GetString("reply")
		root, _ := cmd.Flags().GetString("root")
		anon, _ := cmd.Flags().GetBool("anon")

		useAnon := 0
		if anon {
			useAnon = 1
		}

		if _, err := getClient().PostComment(thread, content, reply, root, useAnon); err != nil {
			return err
		}
		output.PrintSuccess("评论发布成功", jsonOutput)
		return nil
	},
}

var commentDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除评论",
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, _ := cmd.Flags().GetString("thread")
		comment, _ := cmd.Flags().GetString("comment")
		if _, err := getClient().DeleteComment(thread, comment); err != nil {
			return err
		}
		output.PrintSuccess("评论删除成功", jsonOutput)
		return nil
	},
}

var commentLikeCmd = &cobra.Command{
	Use:   "like",
	Short: "点赞评论",
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, _ := cmd.Flags().GetString("thread")
		comment, _ := cmd.Flags().GetString("comment")
		if _, err := getClient().LikeComment(thread, comment); err != nil {
			return err
		}
		output.PrintSuccess("点赞成功", jsonOutput)
		return nil
	},
}

var commentUnlikeCmd = &cobra.Command{
	Use:   "unlike",
	Short: "取消点赞评论",
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, _ := cmd.Flags().GetString("thread")
		comment, _ := cmd.Flags().GetString("comment")
		if _, err := getClient().UnLikeComment(thread, comment); err != nil {
			return err
		}
		output.PrintSuccess("取消点赞成功", jsonOutput)
		return nil
	},
}

func init() {
	commentCmd.AddCommand(commentListCmd)
	commentCmd.AddCommand(commentPostCmd)
	commentCmd.AddCommand(commentDeleteCmd)
	commentCmd.AddCommand(commentLikeCmd)
	commentCmd.AddCommand(commentUnlikeCmd)

	commentListCmd.Flags().String("thread", "", "帖子 ID")
	commentListCmd.MarkFlagRequired("thread")

	commentPostCmd.Flags().String("thread", "", "帖子 ID")
	commentPostCmd.Flags().String("content", "", "评论内容")
	commentPostCmd.Flags().String("reply", "0", "回复的评论 ID")
	commentPostCmd.Flags().String("root", "0", "根评论 ID（楼中楼时使用）")
	commentPostCmd.Flags().Bool("anon", false, "匿名发表")
	commentPostCmd.MarkFlagRequired("thread")
	commentPostCmd.MarkFlagRequired("content")

	commentDeleteCmd.Flags().String("thread", "", "帖子 ID")
	commentDeleteCmd.Flags().String("comment", "", "评论 ID")
	commentDeleteCmd.MarkFlagRequired("thread")
	commentDeleteCmd.MarkFlagRequired("comment")

	commentLikeCmd.Flags().String("thread", "", "帖子 ID")
	commentLikeCmd.Flags().String("comment", "", "评论 ID")
	commentLikeCmd.MarkFlagRequired("thread")
	commentLikeCmd.MarkFlagRequired("comment")

	commentUnlikeCmd.Flags().String("thread", "", "帖子 ID")
	commentUnlikeCmd.Flags().String("comment", "", "评论 ID")
	commentUnlikeCmd.MarkFlagRequired("thread")
	commentUnlikeCmd.MarkFlagRequired("comment")
}

package cmd

import (
	"fmt"

	"github.com/jeanhua/ZanaoCLI/internal/output"
	"github.com/spf13/cobra"
)

var postsCmd = &cobra.Command{
	Use:   "posts",
	Short: "帖子相关操作",
}

var postsListCmd = &cobra.Command{
	Use:   "list",
	Short: "获取帖子列表",
	RunE: func(cmd *cobra.Command, args []string) error {
		fromTime, _ := cmd.Flags().GetString("from")
		posts, err := getClient().GetPost(fromTime)
		if err != nil {
			return err
		}
		output.PrintItems(*posts, jsonOutput)
		return nil
	},
}

var postsHotCmd = &cobra.Command{
	Use:   "hot",
	Short: "获取热门帖子",
	RunE: func(cmd *cobra.Command, args []string) error {
		posts, err := getClient().GetHot()
		if err != nil {
			return err
		}
		output.PrintItems(*posts, jsonOutput)
		return nil
	},
}

var postsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "搜索帖子（当前分类）",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyword, _ := cmd.Flags().GetString("keyword")
		page, _ := cmd.Flags().GetInt("page")
		posts, err := getClient().Search(keyword, page)
		if err != nil {
			return err
		}
		output.PrintItems(*posts, jsonOutput)
		return nil
	},
}

var postsSearchHistoryCmd = &cobra.Command{
	Use:   "search-history",
	Short: "搜索历史帖子",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyword, _ := cmd.Flags().GetString("keyword")
		page, _ := cmd.Flags().GetInt("page")
		rangeTime, _ := cmd.Flags().GetString("range")

		validRanges := map[string]bool{"1d": true, "3d": true, "7d": true, "1m": true, "6m": true, "1y": true}
		if !validRanges[rangeTime] {
			return fmt.Errorf("无效的时间范围 %q，可选值：1d, 3d, 7d, 1m, 6m, 1y", rangeTime)
		}

		posts, err := getClient().SearchHistory(keyword, page, rangeTime)
		if err != nil {
			return err
		}
		output.PrintItems(*posts, jsonOutput)
		return nil
	},
}

var postsLikeCmd = &cobra.Command{
	Use:   "like",
	Short: "点赞帖子",
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, _ := cmd.Flags().GetString("thread")
		if _, err := getClient().LikePost(thread); err != nil {
			return err
		}
		output.PrintSuccess("点赞成功", jsonOutput)
		return nil
	},
}

var postsUnlikeCmd = &cobra.Command{
	Use:   "unlike",
	Short: "取消点赞帖子",
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, _ := cmd.Flags().GetString("thread")
		if _, err := getClient().UnLikePost(thread); err != nil {
			return err
		}
		output.PrintSuccess("取消点赞成功", jsonOutput)
		return nil
	},
}

var postsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "发布新帖子",
	RunE: func(cmd *cobra.Command, args []string) error {
		title, _ := cmd.Flags().GetString("title")
		content, _ := cmd.Flags().GetString("content")
		cate, _ := cmd.Flags().GetString("cate")
		img, _ := cmd.Flags().GetString("img")
		contactPerson, _ := cmd.Flags().GetString("contact-person")
		contactPhone, _ := cmd.Flags().GetString("contact-phone")
		contactQQ, _ := cmd.Flags().GetString("contact-qq")
		contactWX, _ := cmd.Flags().GetString("contact-wx")
		closeComments, _ := cmd.Flags().GetBool("close-comments")

		isCommentClose := 0
		if closeComments {
			isCommentClose = 1
		}

		if _, err := getClient().CreatePost(title, content, cate, img, contactPerson, contactPhone, contactQQ, contactWX, isCommentClose); err != nil {
			return err
		}
		output.PrintSuccess("发帖成功", jsonOutput)
		return nil
	},
}

var postsFinishCmd = &cobra.Command{
	Use:   "finish",
	Short: "将帖子标记为已完成（隐藏发帖人信息）",
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, _ := cmd.Flags().GetString("thread")
		if _, err := getClient().ChangePostStatus(thread, "finish"); err != nil {
			return err
		}
		output.PrintSuccess("帖子已标记为完成", jsonOutput)
		return nil
	},
}

func init() {
	postsCmd.AddCommand(postsListCmd)
	postsCmd.AddCommand(postsHotCmd)
	postsCmd.AddCommand(postsSearchCmd)
	postsCmd.AddCommand(postsSearchHistoryCmd)
	postsCmd.AddCommand(postsLikeCmd)
	postsCmd.AddCommand(postsUnlikeCmd)
	postsCmd.AddCommand(postsCreateCmd)
	postsCmd.AddCommand(postsFinishCmd)

	postsListCmd.Flags().String("from", "0", "分页起始时间戳（0 表示从最新开始，用上一页结果中的 timestamp 获取更早的帖子）")

	postsSearchCmd.Flags().String("keyword", "", "搜索关键词")
	postsSearchCmd.Flags().Int("page", 1, "页码")
	postsSearchCmd.MarkFlagRequired("keyword")

	postsSearchHistoryCmd.Flags().String("keyword", "", "搜索关键词")
	postsSearchHistoryCmd.Flags().Int("page", 1, "页码")
	postsSearchHistoryCmd.Flags().String("range", "1m", "时间范围（1d/3d/7d/1m/6m/1y）")
	postsSearchHistoryCmd.MarkFlagRequired("keyword")

	postsLikeCmd.Flags().String("thread", "", "帖子 ID")
	postsLikeCmd.MarkFlagRequired("thread")

	postsUnlikeCmd.Flags().String("thread", "", "帖子 ID")
	postsUnlikeCmd.MarkFlagRequired("thread")

	postsCreateCmd.Flags().String("title", "", "帖子标题")
	postsCreateCmd.Flags().String("content", "", "帖子内容")
	postsCreateCmd.Flags().String("cate", "", "分类 ID（用 zanao categories 查看）")
	postsCreateCmd.Flags().String("img", "", "图片路径（多个用逗号分隔）")
	postsCreateCmd.Flags().String("contact-person", "", "联系人")
	postsCreateCmd.Flags().String("contact-phone", "", "联系电话")
	postsCreateCmd.Flags().String("contact-qq", "", "联系 QQ")
	postsCreateCmd.Flags().String("contact-wx", "", "联系微信")
	postsCreateCmd.Flags().Bool("close-comments", false, "关闭评论")
	postsCreateCmd.MarkFlagRequired("title")
	postsCreateCmd.MarkFlagRequired("content")
	postsCreateCmd.MarkFlagRequired("cate")

	postsFinishCmd.Flags().String("thread", "", "帖子 ID")
	postsFinishCmd.MarkFlagRequired("thread")
}

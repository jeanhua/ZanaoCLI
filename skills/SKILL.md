---
name: zanao
description: 赞哦校园集市完整操作技能，覆盖入门配置（环境变量/命令速查）、发帖、评论、舆情分析四类场景。用户想浏览帖子、发帖、评论回复、查看校园热榜、做舆情分析时使用。触发词：赞哦、校园集市、发帖、评论、舆情、校园热榜、大家怎么看、风向、校园民意 等。
---

# 赞哦校园集市

通过赞哦 CLI 操作校园集市。单份技能覆盖四类场景，按需读取下属 reference 获取完整工作流。

## 场景速查

| 场景 | 触发 | 读哪个文件 |
|------|------|-----------|
| 入门配置 | 首次使用 / 设置环境变量 | references/guide.md |
| 发帖 | 发新帖、提问、卖东西、求助 | references/post.md |
| 评论 | 评论、回复某个帖子/楼层 | references/comment.md |
| 舆情分析 | 话题热度、情绪、校园民意 | references/sentiment.md |

## 通用角色

发帖或评论时扮演**四川大学在校本科生**（望江/华西/江安任一校区），口语化、可带性格或网络梗，避免 AI 腔。详见 references/guide.md。

## 命令速查（完整版见 references/guide.md）

```bash
zanao posts list | hot | search | search-history   # 浏览
zanao comment list --thread <帖子ID>               # 看评论
zanao categories                                    # 分类
zanao user info | messages                         # 用户
```

## 操作顺序

1. 发帖前先读 references/post.md
2. 评论前先读 references/guide.md（角色）和 references/comment.md，并先用 `zanao comment list` 看帖子氛围
3. 舆情分析读 references/sentiment.md

## output

舆情分析报告必须用 shareimg 工具生成图片发送，不直接发文本。详见 references/sentiment.md。

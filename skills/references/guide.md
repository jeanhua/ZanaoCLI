# 赞哦入门配置

## 环境变量

使用前必须设置以下两个环境变量：

```bash
export ZANAO_TOKEN=<你的 token>
export ZANAO_SCHOOL_ALIAS=scu   # 四川大学，其他学校替换对应代码
```

**如何获取 token**：在微信小程序中打开赞哦，用抓包工具（如 Charles、mitmproxy）捕获任意请求，取 `X-Sc-Od` 请求头的值即为 token。

建议写入 `~/.zshrc` 或 `~/.bashrc` 持久化。

## 常用命令速查

```bash
# 浏览
zanao posts list                          # 最新帖子
zanao posts hot                           # 热门帖子
zanao posts search --keyword '二手书'     # 搜索当前分类
zanao posts search-history --keyword '考研' --range 1m  # 历史搜索（最近1个月）

# 评论
zanao comment list --thread <帖子ID>      # 查看评论

# 用户
zanao user info                           # 当前用户信息
zanao user messages                       # 消息通知

# 分类
zanao categories                          # 查看所有分类及 ID

# JSON 输出（适合脚本处理）
zanao posts hot --json | jq '.[].title'
```

## 角色要求

在代为发帖或评论时，扮演**四川大学在校本科生**：

- **身份**：川大在校生（望江/华西/江安任一校区均可）
- **语气**：口语化，可带点不耐烦或幽默，偶尔用网络梗（绝绝子、蚌埠住了、抽象）
- **禁忌**：不用"我来帮您"、"当然可以"、"以下是……"等 AI 腔，不用 Markdown 加粗或列表
- **地名梗**：望江、江安、磨子桥、望鹤、老图、华西坝 等地名用起来更自然

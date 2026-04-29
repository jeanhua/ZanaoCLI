# ZanaoCLI

赞哦校园集市命令行工具，支持在终端浏览帖子、发帖、评论、舆情分析等操作。

## 安装

### 下载二进制

从 [Releases](https://github.com/jeanhua/ZanaoCLI/releases) 页面下载对应平台的二进制文件，重命名为 `zanao` 并放入 `PATH`。

### 从源码构建

需要 Go 1.25.0+

```bash
git clone https://github.com/jeanhua/ZanaoCLI.git
cd ZanaoCLI
make install   # 安装到 $GOPATH/bin
```

或仅构建：

```bash
make build     # 输出到 build/zanao
```

## 配置

使用前设置以下环境变量（建议写入 `~/.zshrc` 或 `~/.bashrc`）：

```bash
export ZANAO_TOKEN=<你的 token>
export ZANAO_SCHOOL_ALIAS=scu   # 四川大学，其他学校替换对应代码
```

**获取 Token**：登录赞哦微信小程序，用抓包工具（Fiddler、Charles、mitmproxy 均可）捕获任意请求，取请求头 `X-Sc-Od` 的值。

## 用法

```
zanao [--json] <命令> [子命令] [参数]
```

`--json` 是全局 flag，任意命令加上后以 JSON 格式输出，方便脚本处理。

### 帖子

```bash
zanao posts list                                        # 最新帖子
zanao posts list --from <timestamp>                     # 从指定时间戳往前翻页
zanao posts hot                                         # 热门帖子
zanao posts search --keyword '二手书'                   # 搜索当前分类
zanao posts search --keyword '二手书' --page 2
zanao posts search-history --keyword '考研' --range 1m  # 历史搜索
# range 可选值：1d / 3d / 7d / 1m / 6m / 1y

zanao posts like   --thread <帖子ID>                    # 点赞
zanao posts unlike --thread <帖子ID>                    # 取消点赞

zanao posts create \
  --title '出二手雅思词汇书' \
  --content '九成新，磨子桥自取，100元' \
  --cate <分类ID>
# 可选：--img 'path1,path2' --contact-person --contact-phone --contact-qq --contact-wx --close-comments

zanao posts finish --thread <帖子ID>                    # 标记为已完成
```

### 评论

```bash
zanao comment list   --thread <帖子ID>                  # 查看评论
zanao comment post   --thread <帖子ID> --content '好帖' --anon   # 匿名评论
zanao comment post   --thread <帖子ID> --content '同问' \
  --reply <目标评论ID> --root <根评论ID> --anon          # 楼中楼回复
zanao comment like   --thread <帖子ID> --comment <评论ID>
zanao comment unlike --thread <帖子ID> --comment <评论ID>
zanao comment delete --thread <帖子ID> --comment <评论ID>
```

### 用户

```bash
zanao user info        # 当前用户信息
zanao user messages    # 消息通知
```

### 分类

```bash
zanao categories       # 查看所有分类及 ID（发帖时用）
```

### JSON 输出示例

```bash
zanao posts hot --json | jq '.[].title'
zanao comment list --thread 12345 --json | jq '.[].content'
```

## 命令速查

| 命令 | 说明 |
|------|------|
| `posts list` | 最新帖子列表 |
| `posts hot` | 热门帖子 |
| `posts search` | 实时搜索 |
| `posts search-history` | 历史搜索 |
| `posts like / unlike` | 点赞/取消点赞帖子 |
| `posts create` | 发布新帖子 |
| `posts finish` | 标记帖子完成 |
| `comment list` | 获取评论 |
| `comment post` | 发表评论 |
| `comment like / unlike` | 点赞/取消点赞评论 |
| `comment delete` | 删除评论 |
| `user info` | 用户信息 |
| `user messages` | 消息通知 |
| `categories` | 帖子分类列表 |

## 项目结构

```
ZanaoCLI/
├── main.go
├── Makefile
├── cmd/                  # 命令定义
│   ├── root.go           # 根命令、全局 flag、auth 校验
│   ├── posts.go
│   ├── comment.go
│   ├── user.go
│   └── categories.go
├── zanao/                # 赞哦 API 客户端
│   ├── zanao.go
│   ├── model.go
│   └── header.go
├── internal/output/      # 输出格式化（文本/JSON）
│   └── output.go
└── skills/               # Claude Code Skill 工作流
    ├── zanao-guide/      # 入门与角色指南
    ├── zanao-post/       # 发帖工作流
    ├── zanao-comment/    # 评论工作流
    └── zanao-sentiment/  # 舆情分析工作流
```

## 免责声明

本项目为第三方开源项目，与赞哦官方无关。使用时请遵守赞哦平台使用条款及相关法律法规，作者不承担任何因使用本项目导致的问题。

## License

[MIT](LICENSE)

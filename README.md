# AIdea 服务端 - AI 聊天、协作、图像生成

<a href="https://trendshift.io/repositories/855" target="_blank"><img src="https://trendshift.io/api/badge/repositories/855" alt="mylxsw%2Faidea-server | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>

一款集成了主流大语言模型以及绘图模型的 APP 服务端，使用 Golang 开发，代码完全开源。

下载体验地址：

https://aidea.aicode.cc

开源代码：

- 客户端：https://github.com/mylxsw/aidea
- 服务端：https://github.com/mylxsw/aidea-server
- Docker 部署：https://github.com/mylxsw/aidea-docker

## 私有化部署

如果你不想使用托管的云服务，可以自己部署服务端，[部署请看这里](./docs/deploy.md)。

不想自己折腾，可以找我来帮你部署，详情参考 [服务器代部署说明](./docs/deploy-vip.md)。

### 技术交流

- 微信技术交流群：

    <img src="https://github.com/user-attachments/assets/379d0b66-b806-4ed4-ae2e-30fccd9de50e" width="400"/>

    如果无法加入，请添加微信号 `x-prometheus` 为好友，拉你进群。

- 微信公众号

    <img src="https://github.com/mylxsw/aidea-server/assets/2330911/376a3b9f-eacd-45c6-9630-39eb720ba097" width="500" />

## 关于代码

>  目前代码注释、技术文档还比较少，后续有时间会进行补充，敬请见谅。另外以下几点请大家注意，以免造成困扰：
>
> - 代码中 `Room`，`顾问团` 均代表 `数字人`，因项目经过多次改版和迭代，经历了 `房间` -> `顾问团` -> `数字人` 的名称调整
> - 代码中 v1 版本的 `创作岛` 与 v2 版本截然不同，其中 v1 版本服务于 App 1.0.1 及之前版本，从 1.0.2 开始，这部分不再使用，所以就有了
    v2 版本

项目所用的框架

- [Glacier Framework](https://github.com/mylxsw/glacier)： 自研的一款支持依赖注入的模块化的应用开发框架，它以 [go-ioc](https://github.com/mylxsw/go-ioc) 依赖注入容器核心，为 Go 应用开发解决了依赖传递和模块化的问题
- [Eloquent ORM](https://github.com/mylxsw/eloquent) 自研的一款基于代码生成的数据库 ORM 框架，它的设计灵感来源于著名的 PHP 开发框架 Laravel，支持 MySQL 等数据库

代码结构如下

| 目录               | 说明                                                                        |
|------------------|---------------------------------------------------------------------------|
| api              | OpenAI兼容的 API，这里的接口可供第三方支持 OpenAI API 协议的软件直接使用                           |
| server           | 为 AIdea 客户端软件提供的的 API 接口                                                  |
| config           | 配置定义、管理                                                                   |
| migrate          | 数据库迁移文件，SQL 文件                                                            |
| cmd | 程序入口 |
| pkg              | 对外公开的包，其它项目可以直接引用                                                         |
| ⌞ ai             | 不同厂商的 AI 模型接口实现                                                           |
| ⌞ ai/chat        | 聊天模型抽象接口，所有聊天模型都在这里封装为兼容 OpenAI Chat Stream 协议的实现                         |
| ⌞ aliyun         | 阿里云短信、内容安全服务实现                                                            |
| ⌞ dingding       | 钉钉通知机器人                                                                   |
| ⌞ misc           | 部分助手函数                                                                    |
| ⌞ jobs           | 定时任务，用户每日智慧果消耗额度统计等                                                       |
| ⌞ mail           | 邮件发送                                                                      |
| ⌞ proxy          | Socks5 代理实现                                                               |
| ⌞ rate           | 流控实现                                                                      |
| ⌞ redis          | Redis 实例                                                                  |
| ⌞ repo           | 数据模型层，封装了对数据库的操作                                                          |
| ⌞ repo/model     | 数据模型定义，使用了 [mylxsw/eloquent](https://github.com/mylxsw/eloquent)  来创建数据模型 |
| ⌞ service        | Service 层，部分不适合放在 Controller 和 Repo 层的代码，在这里进行封装                          |
| ⌞ sms            | 统一的短信服务封装，对上层业务屏蔽了底层的短信服务商实现                                              |
| ⌞ tencent        | 腾讯语音转文本、短信服务实现                                                            |
| ⌞ token          | JWT Token                                                                 |
| ⌞ uploader       | 基于七牛云存储实现的文件上传下载                                                          |
| ⌞ voice          | 基于七牛云的文本转语音实现，暂时未启用                                                       |
| ⌞ youdao         | 有道翻译服务 API 实现                                                             |
| internal         | 内部包，只有本项目可用 |
| ⌞ queue          | 任务队列实现，所有异步处理的任务都在这里定义                                                    |
| ⌞ queue/consumer | 任务队列消费者                                                                   |
| ⌞ payment        | 在线支付服务实现，如支付宝，Apple                                                       |
| ⌞ coins          | 服务定价、收费策略                                                                 |
| config.yaml      | 配置文件示例                                                                    |
| coins-table.yaml | 价格表配置示例 |
| nginx.conf       | Nginx 配置示例                                                                |
| systemd.service  | Systemd 服务配置示例                                                            |

项目编译：

```bash
go build -o build/debug/aidea-server cmd/main.go
```

## APP 预览图

亮色系

![image](https://github.com/mylxsw/aidea-server/assets/2330911/9c9e878c-67ab-43d6-a9d0-84faf9a6a511)

暗色系

![image](https://github.com/mylxsw/aidea-server/assets/2330911/9e5cc989-4ef5-496b-ab4d-7b9d29793ce3)


## Star History

<a href="https://star-history.com/#mylxsw/aidea-server">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=mylxsw/aidea-server&type=Date&theme=dark" />
    <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=mylxsw/aidea-server&type=Date" />
    <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=mylxsw/aidea-server&type=Date" />
  </picture>
</a>

## License

MIT

Copyright (c) 2023, mylxsw

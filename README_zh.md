# SendUCode-Backend [English](https://github.com/callme-taota/SendUCode)

## 简介
这是一个golang后端项目，帮助您创建设备之间的消息通道。您可以使用此功能在设备（手机到手机、手机到电脑等）之间同步消息。
如果您的设备处于同一网络环境中，则该后端非必需。但如果您的设备不在同一网络环境中，则是必须的。

## 干嘛的
1. 您可以使用此工具同步消息，例如验证码。
2. 您还可以将推送通知从手机转发到PC。

## 相关
1. Windows and MacOS  
   使用electron开发并使用vue作为js框架。  
   [Go to repo]()
2. Android and IOS  
   使用flutter来发展。同时支持Android和IOS。  
   [Go to repo]()

## APIs

| Description  | Methods | Port        | Query           | Header              | Result                   |
|--------------|---------|-------------|-----------------|---------------------|--------------------------|
| 获取消息列表       | GET     | /msg        | limit           | session             | [{detail, device, time}] |
| 发送消息         | POST    | /msg        | message         | session, User-Agent | {msg}                    |
| 确认用户的session | POST    | /user/check | session         |                     | {ok, message, userid}    |
| 新建用户         | POST    | /user       | userid          |                     | {ok, msg, session}       |
| WebSocket 连接 | WS      | /user/ws    | session(Params) | User-Agent          |                          |

## 快速启动
```
git clone https://github.com/callme-taota/SendUCode.git
cd SendUCode
go build main.go
./main.go
```

## 目录结构
```text
├── README.md
├── README_zh.md
├── cache         --- Redis 链接
│   ├── Cache.go  --- Redis 链接功能
│   ├── msg.go    --- 处理缓存消息
│   └── user.go   --- 处理缓存用户
├── conf          --- 本地配置
│   ├── conf.go   --- 配置读取功能
│   └── conf.json --- 配置文件
├── go.mod        --- Go 模块文件
├── go.sum        --- Go sum 文件
├── main.go       --- 入口
├── server        --- 与服务器相关的文件
│   ├── Server.go --- 服务器功能实现文件
│   ├── msg.go    --- 处理消息的实现文件
│   └── user.go   --- 用户操作的实现文件
├── tolog
│   ├── logs      --- 存储日志文件的目录
│   └── tolog.go  --- 处理日志的实现文件
└── utils
├── CreateUserSession.go --- 创建用户会话的实现文件
└── JSONReader.go        --- 读取 JSON 数据的实现文件
```

## Docker
还没写好，即将来临

## 联系我
[我的小站](http://www.callmetaota.fun)
# SendUCode-Backend [简体中文](https://github.com/callme-taota/SendUCode/blob/master/README_zh.md)

## Describe
This is a golang backend project help you to create the message channel between devices. You can use this to synchronous the message between your device(phone to phone, phone to pc,etc... ). 
This is not necessary if your device are in the same network environment. But necessary if your device not in the same network environment.   

## What can you use this for
1. You can use this tool to synchronize messages, such as verification codes.
2. You can also forward push notifications from your mobile to PC.

## Related 
1. Windows and MacOS  
    Using electron to develop and use vue as js framework.   
    [Go to repo]()
2. Android and IOS  
    Using flutter to develop.Both Android and IOS are support .

## APIs 

| Description          | Methods | Port        | Query           | Header              | Result                   |
|----------------------|---------|-------------|-----------------|---------------------|--------------------------|
| Get message list     | GET     | /msg        | limit           | session             | [{detail, device, time}] |
| Send message         | POST    | /msg        | message         | session, User-Agent | {msg}                    |
| CheckUsingSession    | POST    | /user/check | session         |                     | {ok, message, userid}    |
| CreatUser            | POST    | /user       | userid          |                     | {ok, msg, session}       |
| WebSocket connection | WS      | /user/ws    | session(Params) | User-Agent          |                          |

## Quick Start 
```
git clone https://github.com/callme-taota/SendUCode.git
cd SendUCode
go build main.go
./main.go
```

## Directory structure
```text
├── README.md
├── README_zh.md
├── cache         --- Redis Link
│   ├── Cache.go  --- Redis link function
│   ├── msg.go    --- Cache handling messages
│   └── user.go   --- Cache handling user
├── conf          --- Local Config
│   ├── conf.go   --- Config read function
│   └── conf.json --- Config file
├── go.mod        --- Go module files
├── go.sum        --- Go sum files
├── main.go       --- Entry
├── server        --- Server-related files
│   ├── Server.go --- Implementation file for server functionality
│   ├── msg.go    --- Implementation file for handling messages
│   └── user.go   --- Implementation file for user operations
├── tolog
│   ├── logs      --- Directory for storing log files
│   └── tolog.go  --- Implementation file for handling logs.
└── utils
    ├── CreateUserSession.go --- Implementation file for creating user sessions
    └── JSONReader.go        --- Implementation file for reading JSON data
```

## Docker
This dockerfile doesn't contains redis environment , so you have to run redis server yourself .
```
git clone https://github.com/callme-taota/SendUCode.git
cd SendUCode
docker build -t SendUCode .
docker run -e REDIS_HOST=host.docker.internal -p 3003:3003 SendUCode
```
| Environment Option | Description                |
|--------------------|----------------------------|
| REDIS_HOST         | Redis server host          |
| REDIS_PORT         | Redis server port          |
| REDIS_PASSWORD     | Redis server password      |
| REDIS_DB           | Redis server db            |
| SERVER_PORT        | This server listening port |

## Contact
[My website](http://www.callmetaota.fun)
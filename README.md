# go-socket
golang双向通迅demo

运行：
    go run server.go    启动Server端

客户端可以多启动几个，看效果
    go run agent.go
    go run agent.go
    go run agent.go
    go run agent.go
    go run agent.go
    go run agent.go
    go run agent.go
    go run agent.go
    go run agent.go

在客户端命令行，写数据，可以向服务端发送

在Server端命令行，请输入客户端的标识，例如：
2017/10/29 12:54:16 127.0.0.1:50487 clien msg:我要测试
127.0.0.1:50487 这个就是客户端标识
输入之后，就会向这个客户端单独发一条消息


#客户端列表
目前是存在Server端的Map中，使用者可以根据自己的情况，扔到redis,mysql,file等存储


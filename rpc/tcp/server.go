package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (h *HelloService) Say(request string, response *string) error {
	format := time.Now().Format("2006-01-02 15:04:05")
	*response = request + " -- " + format
	return nil
}

func main() {
	// 注册服务名称
	_ = rpc.RegisterName("HelloService", new(HelloService))
	// 监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	for {
		fmt.Println("服务启动...")
		// 监听请求
		accept, err := listen.Accept()
		if err != nil {
			log.Fatalf("Accept Error: %s", err)
		}
		go rpc.ServeConn(accept)
	}
}

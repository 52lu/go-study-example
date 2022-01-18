package main

import (
	"52lu/go-study-example/rpc/jsonrpc/dto"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 建立链接
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		fmt.Println(" rpc.Dial err ", err)
		return
	}
	// 使用json编码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	// 参数
	arg := dto.SumParam{A: 20, B: 80}
	// 结果
	var sum dto.SumRes
	err = client.Call("MathService.Sum", &arg, &sum)
	if err != nil {
		fmt.Println("client.Call err ", err)
		return
	}
	fmt.Printf("res:%+v \n", sum)
}

package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立链接
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println("err ", err)
		return
	}
	// 返回
	var result int
	//  请求方法
	for i := 1; i < 10; i++ {
		err = client.Call("MathService.Multi", i, &result)
		fmt.Printf("i:%v result:%v \n", i, result)
	}
}

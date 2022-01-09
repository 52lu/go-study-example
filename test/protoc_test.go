package test

import (
	"52lu/go-study-example/app/grpc"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestProtocMarshal(t *testing.T) {
	d := &grpc.Data{
		Msg:  "success",
		Code: 1000,
		Data: map[int32]string{
			1:"张三",
			2:"李四",
			3:"赵武",
		},
	}
	marshal, err := proto.Marshal(d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("序列化:",marshal)
	var r grpc.Data
	err = proto.Unmarshal(marshal, &r)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("反序列化:",r)
}

/**
 * @Author Mr.LiuQH
 * @Description encoding/json 测试使用
 * @Date 2021/7/7 10:41 上午
 **/
package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 学校
type School struct {
	Name    string `json:"schoolName"`
	Address string `json:"schoolAddress"`
}
// 学生
type StudentA struct {
	Name string `json:"name"`
	// 匿名字段,而且没有json标签
	School
}
// 序列化-匿名字段 (默认会平铺)
func TestAnonymousFiled(t *testing.T) {
	var student = StudentA{
		Name: "小明",
		School:School{
			Name: "北京大学",
			Address: "北京海淀区",
		},
	}
	jsonByte, _ := json.Marshal(student)
	fmt.Printf("json: %s \n",jsonByte)
}

package main

import (
	"52lu/go-study-example/example/rocketmqdemo/consumer"
	"fmt"
)

type Test struct {

}

func (t Test) Echo(str string)  {
	fmt.Println(str,"echo。。。。")
}

func (t *Test) Hello(str string)  {
	fmt.Println(str,"hello。。。。")
}

func main() {
	consumer.Simple()
}

package main

import (
	"fmt"
)

func main() {
	var a *int
	var b interface{}
	fmt.Printf("a:%v b:%v a == b: %t \n", a, b, a == b)
}

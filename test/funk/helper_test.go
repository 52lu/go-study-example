package funk

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

// 数字型转浮点型
func TestToFloat64(t *testing.T) {
	// int to float64
	d1, _ := funk.ToFloat64(10)
	fmt.Printf("d1 = %v %T \n", d1, d1)
	//@会失败
	d2, err := funk.ToFloat64("10")
	fmt.Printf("d2 = %v %v \n", d2, err)
}

// 相等比较
func TestIsEqual(t *testing.T) {
	// 对比字符串
	fmt.Println("对比字符串:", funk.IsEqual("a", "a"))
	// 对比int
	fmt.Println("对比int:", funk.IsEqual(1, 1))
	// 对比float64
	fmt.Println("对比float64:", funk.IsEqual(float64(1), float64(1)))
	// 对比结构体
	stu1 := Student{Name: "张三", Age: 18}
	stu2 := Student{Name: "张三", Age: 18}
	fmt.Println("对比结构体:", funk.IsEqual(stu1, stu2))
}

// 判断类型是否一样
func TestIsType(t *testing.T) {
	var a, b int8 = 1, 2
	fmt.Println("A: ", funk.IsType(a, b))
	c := 3
	d := "3"
	fmt.Println("B:", funk.IsType(c, d))
}

// 判断是否是array|slice
func TestCollect(t *testing.T) {
	// slice
	a := []int{1, 2, 3}
	// 字符串
	b := "str"
	// 自定义结构体切片
	c := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 18},
	}
	// 数组类型
	d := [2]string{"c", "go"}
	fmt.Println("a:", funk.IsCollection(a))
	fmt.Println("b:", funk.IsCollection(b))
	fmt.Println("c:", funk.IsCollection(c))
	fmt.Println("d:", funk.IsCollection(d))
}

// 返回任一类型的类型切片
func TestSliceOf(t *testing.T) {
	a := []int{10, 20, 30}
	// []int 转成 [][]int
	fmt.Printf("%v %T \n", funk.SliceOf(a), funk.SliceOf(a))
	// string 转成 []string
	fmt.Printf("%v %T \n", funk.SliceOf("go"), funk.SliceOf("go"))
}

// 判断是否为空
func TestEmpty(t *testing.T) {
	// 空结构体
	fmt.Println("空结构体", funk.IsEmpty([]int{}))
	// 空字符串
	fmt.Println("空字符串:", funk.IsEmpty(""))
	// 判断数字0
	fmt.Println("0:", funk.IsEmpty(0))
	// 判断字符串'0'
	fmt.Println("'0':", funk.IsEmpty("0"))
	// nil
	fmt.Println("nil:", funk.IsEmpty(nil))
}

// 生成随机数
func TestRandom(t *testing.T) {
	for i := 1; i <= 3; i++ {
		// 生成任意数字类型
		fmt.Println(funk.RandomInt(1, 100))
	}
}

// 生成随机字符串
func TestRandomString(t *testing.T) {
	for i := 1; i <= 3; i++ {
		// 从默认字符串生成
		fmt.Println("从默认字符串生成:", funk.RandomString(i))
		// 从指的字符串生成
		fmt.Println("从指定字符串生成:", funk.RandomString(i, []rune{'您', '好', '北', '京'}))
	}
}

func TestShard(t *testing.T) {
	tokey := "Hello,Word"
	shard := funk.Shard(tokey, 1, 5, false)
	shard1 := funk.Shard(tokey, 1, 5, true)
	fmt.Println("shard: ", shard)
	fmt.Println("shard1: ", shard1)
	shard2 := funk.Shard(tokey, 2, 5, false)
	shard22 := funk.Shard(tokey, 2, 5, true)
	fmt.Println("shard2: ", shard2)
	fmt.Println("shard22: ", shard22)
}

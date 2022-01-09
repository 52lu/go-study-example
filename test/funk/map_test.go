package funk

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

// 获取map中所有的key
func TestMapKeys(t *testing.T) {
	res := map[string]int{
		"张三": 18,
		"李四": 20,
		"赵武": 25,
	}
	keys := funk.Keys(res)
	fmt.Printf("keys: %#v %T \n", keys, keys)
}

// 获取map中的所有values
func TestMapValues(t *testing.T) {
	res := map[string]int{
		"张三": 18,
		"李四": 20,
		"赵武": 25,
	}
	values := funk.Values(res)
	fmt.Printf("values: %#v %T \n", values, values)
}

// ------------- 数字计算 -------------
// 求最大值
func TestMax(t *testing.T) {
	// 求最大int
	fmt.Println("MaxInt:", funk.MaxInt([]int{30, 10, 8, 11}))
	// 求最大浮点数
	fmt.Println("MaxFloat64:", funk.MaxFloat64([]float64{10.2, 11.0, 8.03}))
	// 求最大字符串
	fmt.Println("MaxString:", funk.MaxString([]string{"a", "d", "c", "b"}))
}

// 求最小值
func TestMin(t *testing.T) {
	// 求最小int
	fmt.Println("MinInt:", funk.MinInt([]int{30, 10, 8, 11}))
	// 求最小浮点数
	fmt.Println("MinFloat64:", funk.MinFloat64([]float64{10.2, 11.0, 8.03}))
	// 求最小字符串
	fmt.Println("MinString:", funk.MinString([]string{"a", "d", "c", "b"}))
}

// 求和
func TestSum(t *testing.T) {
	// 整型
	a := []int{5, 10, 15, 20}
	fmt.Println("int sum:", funk.Sum(a))
	// 浮点型
	b := []float64{5.11, 2.23, 3.31, 0.32}
	fmt.Println("float64 sum:", funk.Sum(b))
}

// 求乘积
func TestProduct(t *testing.T) {
	// 整型
	a := []int{2, 3, 4, 5}
	fmt.Println("int Product:", funk.Product(a))
	// 浮点型
	b := []float64{1.1, 1.2, 1.3, 1.4}
	fmt.Println("float64 Product:", funk.Product(b))
}
func TestNextPermutation(t *testing.T) {
	m := []int{42, 23, 12, 32, 11, 9}
	err := funk.NextPermutation(m)
	if err != nil {
		return
	}
	fmt.Println(m)
}

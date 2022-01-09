package funk

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

func TestInArray(t *testing.T) {
	// string/map/slice/array
	fmt.Println("str1->", funk.Contains("abc", "a"))
	fmt.Println("str2->", funk.Contains("ABC", "A"))
	// []string
	fmt.Println("[]str->", funk.Contains([]string{"a", "b"}, "a"))
	fmt.Println("[]str->", funk.ContainsString([]string{"a", "b"}, "A"))
	// map
	fmt.Println("map->", funk.Contains(map[string]int{"a": 10, "b": 20}, "a"))
}

/*
 * TestExist
 * @Description:
 * @Author: Mr.Liuqh
 * @Param t
 */
func TestExist(t *testing.T) {
	// 判断任意类型
	fmt.Println("str->", funk.Contains([]string{"a", "b"}, "a"))
	// int 类型
	fmt.Println("int->", funk.ContainsInt([]int{1, 2}, 1))
	// in array
	fmt.Println("InInts->", funk.InInts([]int{1, 2}, 1))
	fmt.Println("InStrings1->", funk.InStrings([]string{"a", "b"}, "d"))
	fmt.Println("InStrings2->", funk.InStrings([]string{"a", "b"}, "a"))
}

// 查找元素第一次出现的位置
func TestIndexOf(t *testing.T) {
	strArr := []string{"go", "java", "c", "java"}
	// 具体类型
	fmt.Println("c: ", funk.IndexOfString(strArr, "c"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.IndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.IndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("php: ", funk.IndexOfString(strArr, "php"))
}

// 查找元素最后一次出现的位置
func TestLastOf(t *testing.T) {
	strArr := []string{"go", "java", "c", "java"}
	// 具体类型
	fmt.Println("c: ", funk.LastIndexOfString(strArr, "c"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.LastIndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.LastIndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("php: ", funk.LastIndexOf(strArr, "php"))
}

func TestFind(t *testing.T) {
	// 查找具体类型
	a := []string{"a", "b", "c"}
	res, b := funk.FindString(a, func(s string) bool {
		return s == "b"
	})
	fmt.Printf("findRes: %v b:%t \n ", res, b)
	// 查找值所在的索引位置
	fmt.Println("IndexOf:", funk.IndexOf(a, "b"))
}

// 批量查找，都存在则返回True
func TestEvery(t *testing.T) {
	strArr := []string{"go", "java", "c", "python"}
	fmt.Println("都存在:", funk.Every(strArr, "go", "c"))
	fmt.Println("有一个不存在:", funk.Every(strArr, "php", "go"))
	fmt.Println("都不存在:", funk.Every(strArr, "php", "c++"))
}

// 批量查找，有一则返回true
func TestSome(t *testing.T) {
	strArr := []string{"go", "java", "c", "python"}
	fmt.Println("都存在:", funk.Some(strArr, "go", "c"))
	fmt.Println("至少一个存在:", funk.Some(strArr, "php", "go"))
	fmt.Println("都不存在:", funk.Some(strArr, "php", "c++"))
}

// --------------------------------------------------------------------------
// 填充元素
type Student struct {
	Name string
	Age  int
}

// 填充元素
func TestFill(t *testing.T) {
	// 初始化切片
	var data = make([]int, 3)
	fill, _ := funk.Fill(data, 100)
	fmt.Printf("fill: %v \n", fill)
	// 将所有值设置成2
	input := []int{1, 2, 3}
	result, _ := funk.Fill(input, 2)
	fmt.Printf("result: %v \n", result)
	var structData = make([]Student, 2)
	stuInfo, _ := funk.Fill(structData, Student{Name: "张三", Age: 18})
	fmt.Printf("stuInfo: %v \n", stuInfo)
}

// 访问最后一个元素
func TestLast(t *testing.T) {
	number := []int{20, 40, 60, 80, 100}
	fmt.Println("last: ", funk.Last(number))
}

type cus struct {
	Name string
	Age  int
	Home string
}

// 交集
func TestJoin(t *testing.T) {
	a := []int64{1, 3, 5, 7}
	b := []int64{3, 7}
	// 任意类型切片交集
	join := funk.Join(a, b, funk.InnerJoin)
	fmt.Println("join = ", join)
	// 指定类型取交集
	joinInt64 := funk.JoinInt64(a, b, funk.InnerJoinInt64)
	fmt.Println("joinInt64 = ", joinInt64)
	// 自定义结构体交集
	sliceA := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "南京"},
	}
	sliceB := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "上海"},
	}
	res := funk.Join(sliceA, sliceB, funk.InnerJoin)
	fmt.Println("自定义结构体: ", res)
}

// 取差集
func TestDiffSlice(t *testing.T) {
	a := []int64{1, 3, 5, 7}
	b := []int64{3, 7, 10}
	// 任意类型切片交集
	join := funk.Join(a, b, funk.OuterJoin)
	fmt.Println("OuterJoin = ", join)
	// 指定类型取交集
	joinInt64 := funk.JoinInt64(a, b, funk.OuterJoinInt64)
	fmt.Println("joinInt64 = ", joinInt64)
	// 自定义结构体交集
	sliceA := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "南京"},
	}
	sliceB := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "上海"},
	}
	res := funk.Join(sliceA, sliceB, funk.OuterJoin)
	fmt.Println("自定义结构体: ", res)
}

func TestLeftAndRightJoin(t *testing.T) {
	a := []int64{10, 20, 30}
	b := []int64{30, 40}
	// 取出只在a，不在b的元素
	leftJoin := funk.Join(a, b, funk.LeftJoin)
	fmt.Println("只在a切片的元素: ", leftJoin)
	// 取出只在b，不在a的元素
	rightJoin := funk.Join(a, b, funk.RightJoin)
	fmt.Println("只在b切片的元素: ", rightJoin)
}

func TestDifferent(t *testing.T) {
	// 处理任意类型
	one := []interface{}{1, "go", 3.2, []int8{10}}
	two := []interface{}{2, "go", 3.2, []int{20}}
	oneRes, twoRes := funk.Difference(one, two)
	fmt.Println("oneRes: ", oneRes)
	fmt.Println("twoRes:  ", twoRes)

	// 只处理具体类型
	str1 := []string{"go", "php", "java"}
	str2 := []string{"c", "python", "java"}
	res, res1 := funk.DifferenceString(str1, str2)
	fmt.Println("res: ", res)
	fmt.Println("res1: ", res1)
}

// 获取第一个元素 或 最后一个元素
func TestLastOrFirst(t *testing.T) {
	number := []int{10, 30, 12, 23}
	// 获取第一个元素
	fmt.Println("Head: ", funk.Head(number))
	// 获取最后一个元素
	fmt.Println("Last: ", funk.Last(number))
}

// 除去第一个 或者 最后一个
func TestDelLastOrFirst(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	// 除去第一个，返回剩余元素
	fmt.Println(funk.Tail(a))
	// 除去最后一个，返回剩余元素
	fmt.Println(funk.Initial(a))
}

// 打乱数组
func TestShuffle(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	// 打乱多次
	for i := 1; i <= 3; i++ {
		fmt.Println(fmt.Sprintf("第%v次打乱a", i), funk.Shuffle(a))
	}
}

// 反转切片
func TestReverse(t *testing.T) {
	fmt.Println("ReverseInt:", funk.ReverseInt([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println("Reverse,任意类型:", funk.Reverse([]interface{}{1, "2", 3.02, 4, "5", 6}))
	fmt.Println("Reverse,Str:", funk.ReverseStrings([]string{"a", "b", "c", "d"}))
}

/*
 * TestCompact
 * @Description: 压缩
 * @Author: LiuQHui
 * @Param t
 */
func TestCompact(t *testing.T) {
	d := []interface{}{"Hello", ",", "Work", 1, 2, 3, nil}
	compact := funk.Compact(d)
	fmt.Printf("%v %T \n", compact, compact)
}

// 去重
func TestUniq(t *testing.T) {
	a := []int64{1, 2, 3, 4, 3, 2, 1}
	// 过滤整型类型
	fmt.Println("Uniq:", funk.Uniq(a))
	fmt.Println("UniqInt64:", funk.UniqInt64(a))
	b := []string{"php", "go", "c", "go"}
	// 过滤字符串类型
	fmt.Println("UniqString:", funk.UniqString(b))
}

// 删除制定元素
func TestWithOut(t *testing.T) {
	// 删除具体元素
	b := []int{10, 20, 30, 40}
	without := funk.Without(b, 30)
	fmt.Println("删除30:", without)
	// 删除自定义结构体元素
	stuList := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 22},
		{Name: "范五", Age: 24},
	}
	res := funk.Without(stuList, Student{Name: "李四", Age: 22})
	fmt.Println("删除李四:", res)
}

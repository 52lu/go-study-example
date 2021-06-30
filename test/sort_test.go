/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/6/16 7:26 下午
 **/
package test

import (
	"fmt"
	"sort"
	"testing"
)

// 排序
func TestStringSort(t *testing.T) {
	stringSlice := sort.StringSlice{"d", "a", "c", "e", "b"}
	fmt.Printf("排序前:%v\n", stringSlice)
	stringSlice.Sort()
	fmt.Printf("排序后:%v\n", stringSlice)
}

// 查找
func TestSearch(t *testing.T) {
	// 1.定义类型
	intSlice := sort.IntSlice{10, 40, 6, 9, 23, 72, 89}
	// 2.排序后查找
	intSlice.Sort()
	// 3.查询可能存在的索引
	searchVal := 40
	searchIndex := intSlice.Search(searchVal)
	// 4.判断是否相等
	if searchIndex > len(intSlice) {
		fmt.Printf("未找到：%v\n", searchVal)
	} else {
		if searchVal == intSlice[searchIndex] {
			fmt.Printf("找到: %v 索引为: %v \n", searchVal, searchIndex)
		} else {
			fmt.Printf("未找到：%v\n", searchVal)
		}
	}

	// 查找不存在情况
	searchVal2 := 90
	searchIndex2 := intSlice.Search(searchVal2)
	if searchIndex2 >= len(intSlice) {
		fmt.Printf("未找到：%v\n", searchVal2)
	} else {
		if searchVal2 == intSlice[searchIndex2] {
			fmt.Printf("找到: %v 索引为: %v \n", searchVal2, searchIndex2)
		} else {
			fmt.Printf("未找到：%v\n", searchVal2)
		}
	}
	// 逆序
	sort.Sort(sort.Reverse(intSlice))
	fmt.Printf("逆序：%v\n", intSlice)

}

// 测试使用
func TestInts(t *testing.T) {
	a := []int{23, 12, 9, 2, 12, 3, 89}
	// 调用排序
	sort.Ints(a)
	// 检测是否已经排序
	if sorted := sort.IntsAreSorted(a); sorted {
		fmt.Printf("排序成功: %v\n", a)
	} else {
		fmt.Printf("排序失败: %v\n", a)
	}
	// 查找元素
	fmt.Printf("查找元素(存在): %v 索引:%v \n", 12, sort.SearchInts(a, 12))
	fmt.Printf("查找元素(不存在): %v 索引:%v \n", 22, sort.SearchInts(a, 22))
}

// 浮点型切片排序
func TestFloat64s(t *testing.T) {
	f := []float64{20.2, 13.8, 28.12, 5.23}
	// 排序
	sort.Float64s(f)
	// 检测是否已经排序
	if sorted := sort.Float64sAreSorted(f); sorted {
		fmt.Printf("排序成功: %v\n", f)
	} else {
		fmt.Printf("排序失败: %v\n", f)
	}
	// 查找元素
	searchVal1 := 13.8
	searchVal2 := 29.23
	fmt.Printf("查找元素(存在): %v 索引:%v \n", searchVal1, sort.SearchFloat64s(f, searchVal1))
	fmt.Printf("查找元素(不存在): %v 索引:%v \n", searchVal2, sort.SearchFloat64s(f, searchVal2))
}

// 字符串型切片排序
func TestSlice(t *testing.T) {
	s := []string{"d", "a", "c", "e", "b"}
	// 排序
	sort.Strings(s)
	// 检测是否已经排序
	if sorted := sort.StringsAreSorted(s); sorted {
		fmt.Printf("排序成功: %v\n", s)
	} else {
		fmt.Printf("排序失败: %v\n", s)
	}
	// 查找元素
	searchVal1 := "e"
	searchVal2 := "f"
	fmt.Printf("查找元素(存在): %v 索引:%v \n", searchVal1, sort.SearchStrings(s, searchVal1))
	fmt.Printf("查找元素(不存在): %v 索引:%v \n", searchVal2, sort.SearchStrings(s, searchVal2))
}

// 测试降序使用
func TestReverseSort(t *testing.T) {
	// 使用sort包中的数据类型
	intSlice := sort.IntSlice{23, 15, 9, 2, 12, 3, 89}
	// 互换j,i
	reverse := sort.Reverse(intSlice)
	// 排序
	sort.Sort(reverse)
	fmt.Printf("排序结果: %v\n", intSlice)
}
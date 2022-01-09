package funk

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

type User struct {
	Name string
	Home struct {
		City string
	}
}

func TestGetToSlice(t *testing.T) {
	userList := []User{
		{
			Name: "张三",
			Home: struct {
				City string
			}{"北京"},
		},
		{
			Name: "小明",
			Home: struct {
				City string
			}{"南京"},
		},
	}
	// 取一层
	names := funk.Get(userList, "Name")
	fmt.Println("names:", names)
	// 取其他层
	homes := funk.Get(userList, "Home.City")
	fmt.Println("homes:", homes)
}

// 遍历切片
func TestForeachSlice(t *testing.T) {
	// 从左边遍历
	var leftRes []int
	funk.ForEach([]int{1, 2, 3, 4}, func(x int) {
		leftRes = append(leftRes, x*2)
	})
	fmt.Println("ForEach:", leftRes)
	// 从右边遍历
	var rightRes []int
	funk.ForEachRight([]int{1, 2, 3, 4}, func(x int) {
		rightRes = append(rightRes, x*2)
	})
	fmt.Println("ForEachRight:", rightRes)
}

// 三元运算
func TestShortIf(t *testing.T) {
	fmt.Println("10 > 5 :", funk.ShortIf(10 > 5, 10, 5))
	fmt.Println("10.0 == 10 : ", funk.ShortIf(10.0 == 10, "yes", "no"))
	fmt.Println("'a' == 'b' : ", funk.ShortIf('a' == 'b', "equal chars", "unequal chars"))
}

// 判断是否属于子集
func TestSubset(t *testing.T) {
	// 判断基础切片
	fmt.Println("是否属于子集:", funk.Subset([]int{1}, []int{1, 2}))
	// 判断自定义结构体切片
	subStu1 := []Student{{Name: "张三", Age: 18}}
	subStu2 := []Student{{Name: "张三", Age: 22}}
	allStu := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 22},
	}
	fmt.Println("subStu1: ", funk.Subset(subStu1, allStu))
	fmt.Println("subStu2: ", funk.Subset(subStu2, allStu))
	// 判断空切片是否属于另一切片子集
	fmt.Println("判断空集：", funk.Subset([]int{}, []int{1, 2}))
}

// 分组
func TestChunk(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	// 分组(每组2个)
	fmt.Println(funk.Chunk(a, 2))
	// 分组自定义结构体切片 (每组2个)
	stuList := []Student{
		{Name: "张三", Age: 18},
		{Name: "小明", Age: 20},
		{Name: "李四", Age: 22},
		{Name: "赵武", Age: 32},
		{Name: "小英", Age: 19},
	}
	fmt.Println(funk.Chunk(stuList, 2))
}

type Book struct {
	Id   int
	Name string
}

// 结构体转成Map
func TestToMap(t *testing.T) {
	bookList := []Book{
		{Id: 1, Name: "西游记"},
		{Id: 2, Name: "水浒传"},
		{Id: 3, Name: "三国演义"},
	}
	// 转成以Id为Key的Map
	fmt.Println("结果:", funk.ToMap(bookList, "Id"))
}

func TestMap(t *testing.T) {
	// 将切片最为map的key
	r := funk.Map([]int{1, 2, 3, 4}, func(x int) (int, string) {
		return x, "go"
	})
	fmt.Println("r=", r)
}

// 把二维切片转成一维切片
func TestFlatMap(t *testing.T) {
	r := funk.FlatMap([][]int{{1}, {2}, {3}, {4}}, func(x []int) []int {
		return x
	})
	fmt.Printf("%#v\n", r)
}

type School struct {
	Name    string `json:"nc"`
	Address string `json:"address"`
}

func TestPruneByTag(t *testing.T) {
	b := School{
		Name:    "北京大学",
		Address: "北京",
	}
	tag, err := funk.PruneByTag(b, []string{"Name"}, "nc")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", tag)
}

package test

import (
	"52lu/go-study-example/example/collyDemo"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"net/http"
	"testing"
)

func TestCollyDemo(t *testing.T) {
	err := collyDemo.RunDemo()
	if err != nil {
		t.Error(err)
	}
}

func TestDouBan(t *testing.T) {
	err := collyDemo.DouBanBook()
	if err != nil {
		t.Error(err)
	}
}

func TestJiJIn(t *testing.T) {
	url := "https://www.morningstar.cn/quickrank/rqfii.aspx"
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"),
		colly.Debugger(&debug.LogDebugger{}),
		//colly.Async(true),
	)
	cookieList := []*http.Cookie{
		&http.Cookie{
			Name:     "authWeb",
			Value:    "AFE9FB65A838E022E595BC294AD98AE8B78C2DE6FE4484976A8E032513B0E4FAEC101AF2D212BE94AD36A029DAAF436615769E72A0966B8A97C556FB741ADAC6C4F706B0C427DFBB53F7A6E176E7E4D21E525BE38877C3657721969286FFCA34168496E2087E8C11692A7D0DEABD72818CAD8038",
			Path:     "/",
			Domain:   ".morningstar.cn",
			Secure:   true,
			HttpOnly: true,
		},
	}
	err := collector.SetCookies(url, cookieList)
	if err != nil {
		t.Errorf("set cookie error:%s", err)
		return
	}
	collector.OnError(func(response *colly.Response, err error) {
		t.Errorf("req error:%s", err)
		return
	})
	collector.OnHTML("#ctl00_cphMain_gridResult", func(element *colly.HTMLElement) {
		fmt.Println("tbody:", element.ChildText("tbody"))
	})
	err = collector.Visit(url)
	collector.OnError(func(response *colly.Response, err error) {
		t.Errorf(" error:%s", err)
		return
	})
	//collector.Wait()
}

// 返回当前元素的属性
func TestUseAttr(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("div[class='nav-logo'] > a", func(element *colly.HTMLElement) {
		// 定位到div[class='nav-logo'] > a标签元素
		fmt.Printf("href:%v\n",element.Attr("href"))
	})
	_ = collector.Visit("https://book.douban.com/tag/小说")
}

// 测试使用ChildAttr和ChildAttrs
func TestChildAttrMethod(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("OnError",err)
	})
	//
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// 获取第一个子元素(div)的class属性
		fmt.Printf("ChildAttr:%v\n",element.ChildAttr("div","class"))
		// 获取所有子元素(div)的class属性
		fmt.Printf("ChildAttrs:%v\n",element.ChildAttrs("div","class"))
	})
	err := collector.Visit("https://liuqh.icu/a.html")
	if err != nil {
		fmt.Println("err",err)
	}
}


// 测试使用ChildText和ChildTexts
func TestChildTextMethod(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("OnError",err)
	})
	//
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// 获取第一个子元素(div)的class属性
		fmt.Printf("ChildText:%v\n",element.ChildText("div"))
		// 获取所有子元素(div)的class属性
		fmt.Printf("ChildTexts:%v\n",element.ChildTexts("div"))
	})
	err := collector.Visit("https://liuqh.icu/a.html")
	if err != nil {
		fmt.Println("err",err)
	}
}
// 遍历
func TestForeach(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("ul[class='demo']", func(element *colly.HTMLElement) {
		element.ForEach("li", func(_ int, el *colly.HTMLElement) {
			name := el.ChildText("span[class='name']")
			age := el.ChildText("span[class='age']")
			home := el.ChildText("span[class='home']")
			fmt.Printf("姓名: %s  年龄:%s 住址: %s \n",name,age,home)
		})
	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}

// 定义结构体
type Book struct {
	Name string `selector:"span.title"`
	Link string `selector:"span > a" attr:"href"`
	Author string `selector:"span.autor"`
	Reviews []string `selector:"ul.category > li"`
	Price string `selector:"span.price"`
}

func TestUnmarshal(t *testing.T) {
	// 声明结构体
	var book Book
	collector := colly.NewCollector()
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		err := element.Unmarshal(&book)
		if err != nil {
			fmt.Println("解析失败:",err)
		}
		fmt.Printf("结果:%+v\n",book)
	})
	_ = collector.Visit("https://liuqh.icu/a.html")
}

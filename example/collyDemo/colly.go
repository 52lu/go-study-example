package collyDemo

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func RunDemo() error {
	// 创建 Collector 对象
	collector := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"))
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("在请求之前调用")
	})
	// 请求期间发生错误,则调用
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("请求错误:",err)
	})
	// 收到响应后调用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("url:",response.Request.URL)
	})
	//OnResponse如果收到的内容是HTML ,则在之后调用
	collector.OnHTML("#position_shares table", func(element *colly.HTMLElement) {
		// 遍历table
		element.ForEach("table tr", func(_ int, el *colly.HTMLElement) {
			name := el.ChildText("td:nth-of-type(1)")
			percentage := el.ChildText("td:nth-of-type(2)")
			fmt.Printf("名称:%v 仓位占比:%v \n",name,percentage)
		})

	})
	return collector.Visit("https://fund.eastmoney.com/481010.html")
}
package main

import (
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"strings"
)

var URL = "http://tonkiang.us/"

func main() {
	var json map[string]string
	//router := gin.New()
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0"),
	)

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("Content-Type", "application/x-www-form-urlencoded")

		log.Println("正在请求")
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Fatal("请求错误", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == http.StatusOK {
			log.Println("响应成功")
		} else {
			log.Fatal("响应失败")
		}
	})

	c.OnHTML(".tables .result", func(element *colly.HTMLElement) {
		//element.ForEach(".channel", func(i int, element *colly.HTMLElement) {
		//	text := element.ChildText("div:first-child")
		//	println("标题：", text)
		//})
		//
		//element.ForEach(".m3u8 > table > tbody > tr > td", func(i int, element *colly.HTMLElement) {
		//	if strings.Contains(element.Text, "m3u8") {
		//		println("链接：", element.Text)
		//	}
		//})
		//
		//element.ForEach("i", func(i int, element *colly.HTMLElement) {
		//	println("信息：", trimSpace(element.Text))
		//})

		json = map[string]string{
			"result": "hello world",
		}
	})

	c.OnScraped(func(response *colly.Response) {
		println("所有匹配已完成")
		println(json)
	})

	requestData := map[string]string{
		"search": "cctv1",
	}

	err := c.Post(URL, requestData)
	if err != nil {
		println("err", err.Error())
		return
	}
}

func trimSpace(str string) string {
	//str = strings.ReplaceAll(str, " ", "")

	return strings.ReplaceAll(str, "\n", "")
	//return str
}

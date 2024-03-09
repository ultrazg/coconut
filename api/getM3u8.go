package api

import (
	C "coconut/constant"
	M "coconut/model"
	U "coconut/util"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"time"
)

func GetM3u8(keyword string) (int, float64, string, *[]M.Records, error) {
	var records []M.Records
	startTime := time.Now()

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0"),
	)

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("Content-Type", "application/x-www-form-urlencoded")

		log.Println("fetching")
	})

	c.OnError(func(response *colly.Response, err error) {
		panic("bad request" + err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == http.StatusOK {
			log.Println("successful response")
		} else {
			log.Fatal("response failure")
		}
	})

	c.OnHTML(".tables", func(element *colly.HTMLElement) {
		element.ForEach(".result", func(i int, element *colly.HTMLElement) {
			title := element.ChildText(".channel > a > div:first-child")
			url := element.ChildText(".m3u8 > table > tbody > tr > td")
			info := U.TrimSpace(element.ChildText("i"))

			records = append(records, M.Records{
				Title: title,
				Url:   url,
				Info:  info,
			})
		})
	})

	c.OnScraped(func(response *colly.Response) {
		log.Println("all matches are complete")
	})

	requestData := map[string]string{
		"search": keyword,
	}

	err := c.Post(C.Url, requestData)
	if err != nil {
		log.Fatalln("err", err.Error())

		return 0, 0, "Error", nil, err
	}

	duration := time.Since(startTime)
	return len(records), duration.Seconds(), "OK", &records, nil
}

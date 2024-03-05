package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	var url string = "https://www.jumia.com.ng/xiaomi-curved-gaming-monitor-30-250366429.html"
	scrape(url)
}

func scrape(url string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML("#jm", func(h *colly.HTMLElement) {
		fmt.Println(h)
	})

	c.Visit(url)
}

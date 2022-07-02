package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("meta[property='og:image']", func(e *colly.HTMLElement) {
		fmt.Println("og:image -> ", e.Attr("content"))
	})

	c.Visit("https://www.tsuyukimakoto.com/blog/2022/06/29/summer/")
}
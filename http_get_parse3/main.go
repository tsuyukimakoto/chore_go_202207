package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	resp, err := http.Get("https://www.tsuyukimakoto.com/blog/2022/06/29/summer/")
	// resp, err := http.Get("https://www.tsuyukimakoto.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if val, exists := doc.Find("meta[property='og:image']").Attr("content"); exists {
		fmt.Println("og:image -> ", val)
	} else {
		fmt.Println("No og:image")
	}
}
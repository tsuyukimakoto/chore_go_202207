package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func searchOgImage(targetUrl string) (url string, err error) {
	resp, err := http.Get(targetUrl)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	for {

		if tokenType := tokenizer.Next(); tokenType == html.ErrorToken {
			if err := tokenizer.Err(); err == io.EOF {
				fmt.Println("Finish!")
				break
			}
			log.Fatalf("Error tokenize HTML: %v\n", tokenizer.Err())
		}

		if tagName, hasAttr := tokenizer.TagName(); string(tagName) == "meta" && hasAttr {
			// fmt.Println("Tag: ", string(tagName))
			og_image_found := false
			for {
				key, val, hasMore := tokenizer.TagAttr()
				if string(key) == "property" && string(val) == "og:image" {
					og_image_found = true
				}
				if og_image_found && string(key) == "content" {
					return string(val), nil
				}
				// fmt.Println("  Attribute: ", string(key), string(val))
				if !hasMore {
					break
				}
			}
		}
	}
	return "", errors.New("no og:img")
} 

func main() {
	url, err := searchOgImage("https://www.tsuyukimakoto.com/blog/2022/06/29/summer/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("og:img -> ", url)

// <meta property="og:image" content="https://www.tsuyukimakoto.com/static/blog/2022/0629/bridge.jpg"/>
}
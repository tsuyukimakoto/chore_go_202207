package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://www.tsuyukimakoto.com/blog/2022/06/29/summer/")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				fmt.Println("Finish!")
				break
			}
			log.Fatalf("Error tokenize HTML: %v\n", tokenizer.Err())
		}
		
		if tagName, hasAttr := tokenizer.TagName(); hasAttr {
			fmt.Println("Tag: ", string(tagName))
			for {
				key, val, hasMore := tokenizer.TagAttr()
				fmt.Println("  Attribute: ", string(key), string(val))
				if !hasMore {
					break
				}
			}
		}
	}

 		// case html.TextToken:
		// fmt.Println("TextToken: ", tokenizer.Token().Data)
		// case html.EndTagToken:
		// case html.SelfClosingTagToken:
		// case html.CommentToken:
		// case html.DoctypeToken:

// <meta property="og:image" content="https://www.tsuyukimakoto.com/static/blog/2022/0629/bridge.jpg"/>
}
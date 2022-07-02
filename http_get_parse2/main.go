package main

import (
	"fmt"
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

	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Parse Error: %v\n", err)
	}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Println("Tag -> ", n.Data)
			for _, attr := range n.Attr {
				fmt.Println("  Attr -> ", attr.Key, attr.Val)
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}
	traverse(root)
}
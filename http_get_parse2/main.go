package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		if n.Type == html.ElementNode && n.Data == "meta" {
			// fmt.Println("Tag -> ", n.Data)
			fmt.Print("Found!")
			og_image_found := false
			for _, attr := range n.Attr {
				if attr.Key == "property" && attr.Val ==  "og:image" {
					og_image_found = true
				}
				if og_image_found && attr.Key == "content" {
					fmt.Println("og:image -> ", attr.Val)
					os.Exit(0)
				}
				// fmt.Println("  Attr -> ", attr.Key, attr.Val)
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}
	traverse(root)
}
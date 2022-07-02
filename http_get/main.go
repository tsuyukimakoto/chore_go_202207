package main

import (
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

func main() {
	url := "https://www.tsuyukimakoto.com/"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error: NewRequest ", err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error: client.Do ", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	// fmt.Println("Body", body) ← byte[]なので数値が出力される
	fmt.Println(*(*string)(unsafe.Pointer(&body)))
	// fmt.Println(string(body)) ← ↑よりは遅いっぽいけど十分かな？Printfはかなり遅そう
	
}
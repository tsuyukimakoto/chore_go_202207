package main

import (
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

func main() {
	client := new(http.Client)
	resp, err := client.Get("https://www.tsuyukimakoto.com/")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(*(*string)(unsafe.Pointer(&body)))
}
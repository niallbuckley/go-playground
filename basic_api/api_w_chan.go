package main

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"
)

func makeApiReq(url string, channel chan[]byte){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	byteSl, err := io.ReadAll(resp.Body)

	channel <- byteSl
}

func main(){
	url:="https://jsonplaceholder.typicode.com/posts/1"

	respChan := make(chan[]byte, 1)
	go makeApiReq(url, respChan)

	// map with string as key and interface as val
	// value is string because JSON keys MUST be strings.
	var data map[string]interface{}

	buff  := <- respChan
	_ = json.Unmarshal(buff, &data)

	fmt.Println(data["userId"])
}

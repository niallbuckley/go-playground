package main

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"
)



func main(){
	url:="https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	buff, err := io.ReadAll(resp.Body)

	var data map[string]interface{}
	_ = json.Unmarshal(buff, &data)

	fmt.Println(data["userId"])
}

package main

import "fmt"

func getPointer() *int {
    x := 10
    return &x
}

func main() {
    p := getPointer()
    fmt.Println(*p)
}


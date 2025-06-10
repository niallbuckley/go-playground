package main

import (
	"fmt"
)

func testerrA() (string, error) {
    return "", fmt.Errorf("testerr failed on A")
}

func testerrB() (string, error) {
    return "", fmt.Errorf("testerr failed on B")
}

func main() {
    fmt.Println("START")


    a, err := testerrA()
    b, err := testerrB()
    
    fmt.Println("a, b: ", a, b)
    fmt.Println("err: ", err)

    fmt.Println("END")
}

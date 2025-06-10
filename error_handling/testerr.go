package main

import (
	"fmt"
)

func testerr() (string, error) {
    return "", fmt.Errorf("testerr failed")
}

func main() {
    fmt.Println("START")


    var err error
    var b string
    for attempt := 0; attempt < 2; attempt++ {
// Becuase err is reclared in the for loop it only exists in the scope of the for loop it will not ne printed.
        a, err := testerr()
        if err != nil {
            b = a
        }
    }
    fmt.Println("b: ", b)
    fmt.Println("err: ", err)

    fmt.Println("END")
}

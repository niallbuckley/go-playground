package main 

import (
	"fmt"
)

var supervar = "global"

func main() {
    fmt.Println("START")

    var supervar = "local"
    if true {
        var supervar = "block"
        fmt.Printf("supervar inside block = %s\n", supervar)
    }

    fmt.Printf("supervar inside main = %s\n", supervar)

    fmt.Println("END")
}

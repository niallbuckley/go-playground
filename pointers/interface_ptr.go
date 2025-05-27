package main

import "fmt"


func changeVal(i interface{}) {
    i = 10
}

func changeValPtr(i *interface{}) {
    *i = 10
}

func main() {
    var i interface{} = 5
    changeVal(i)
    fmt.Println(i)
    changeValPtr(&i)
    fmt.Println(i)
}


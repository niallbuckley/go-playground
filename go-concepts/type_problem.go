package main

import "fmt"

type MyStruct struct {
    Name string
}

func (m MyStruct) String() string {
    return fmt.Sprintf("My name is %s", m.Name)
}

func main() {
    var s fmt.Stringer
    if s == nil {
        fmt.Println("nil interface")
    }
    s = MyStruct{Name: "Go"}
    fmt.Println(s)
}

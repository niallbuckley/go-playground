package main

import "fmt"

type Person struct {
    Name string
}

func changeName(p *Person) {
    p.Name = "Alice"
}

func main() {
    p := Person{Name: "Bob"}
    changeName(&p)
    fmt.Println(p.Name)
}


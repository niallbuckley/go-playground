package main 
import "fmt"

func foo(p *int) {
    *p = 10
}

func main() {
    x := 5
    foo(&x)
    fmt.Println(x)

}

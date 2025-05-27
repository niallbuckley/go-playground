package main

import "fmt"

func appendValue(slice []int) {
    slice = append(slice, 10)
}

func appendPointer(slice *[]int) {
    *slice = append(*slice, 20)
}

func modify(slice []int) {
    slice[0] = 100
}

func modifyPtr(slice *[]int) {
    (*slice)[0] = 200
}

func main() {
    s := []int{1, 2, 3}

    appendValue(s)
    fmt.Println("After appendValue:", s)

    appendPointer(&s)
    fmt.Println("After appendPointer:", s)


    s = []int{1, 2, 3}
    modify(s)
    fmt.Println(s)
    modifyPtr(&s)
    fmt.Println(s)

}

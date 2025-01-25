package main

import "fmt"

// Recovery function to handle panics
func handlePanic() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
}

// Safe division function
func safeDivision(a, b int) int {
    defer handlePanic() // Call the named function to recover from panics
    return a / b // This can panic if b == 0
}

func main() {
    fmt.Println(safeDivision(10, 2)) // Normal division
    fmt.Println(safeDivision(10, 0)) // This will trigger a panic, but it will be recovered
}

package main

import "fmt"

func sendValues(ch chan int) {
    ch <- 1 // Doesn't block since the buffer has capacity
    ch <- 2 // Doesn't block since the buffer has capacity
}

func receiveValues(ch chan int) {
    fmt.Println(<-ch) // Receives the first value
    fmt.Println(<-ch) // Receives the second value
}

func main() {
    ch := make(chan int, 2) // Buffered channel with a capacity of 2

    sendValues(ch)   // Send values to the channel
    receiveValues(ch) // Receive values from the channel
}


package main

import "fmt"

func unbuffered(myChan chan int){
	myChan <- 1 // This will block until the main goroutine receives the value.
  
}

func main() {
    ch := make(chan int)
    
    go unbuffered(ch) 

    fmt.Println(<-ch) // Receiver ready to receive the value.
}

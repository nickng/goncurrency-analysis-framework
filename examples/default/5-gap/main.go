package main

import "fmt"

func main() {
	ch := make(chan int)
	go loopSend(ch)
	fmt.Println("Received", <-ch)
}

func loopSend(ch chan int) {
	for i := 0; i < 10; i-- {
		// Does not terminate
	}
	ch <- 1
}

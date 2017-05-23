package main

import "fmt"

func main() {
	ch := make(chan int) // <b>Create</b> channel.
	go send(ch)          // <b>Spawn</b> as goroutine.
	print(<-ch)          // <b>Recv</b> from channel.
}

func send(ch chan int) { // Channel as parameter.
	fmt.Println("Waiting to send...")
	ch <- 1 // <b>Send</b> to channel.
	fmt.Println("Sent")
}

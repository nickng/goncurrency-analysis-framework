package main

import "fmt"

//import _ "net" // Load "net" package

func main() {
	ch := make(chan int) // <b>Create</b> channel.
	send(ch)             // <s>Spawn</s> as goroutine.
	print(<-ch)          // <b>Recv</b> from channel.
}

func send(ch chan int) { // Channel as parameter.
	fmt.Println("Waiting to send...")
	ch <- 1 // Send to channel.
	fmt.Println("Sent")
}

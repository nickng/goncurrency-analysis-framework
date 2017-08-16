package main

// <b>Eventual reception</b>
// Message in buffer will eventually be received.

func main() {
	ch := make(chan int, 1)
	ch <- 1
}

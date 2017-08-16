package main

// <b>Channel Safety</b>
// Closed channel cannot be closed twice
// Cannot send on closed channel

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	<-ch
	close(ch)
	// <-ch
	// ch <- 1 // Send on closed channel
}

package main

// <b>Channel Safety</b>
// Closed channel cannot be closed twice

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	<-ch
	close(ch)
}

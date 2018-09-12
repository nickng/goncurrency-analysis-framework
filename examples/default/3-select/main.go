package main

func recv(ch chan int) {
	<-ch
}
func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	go recv(a)
	select {
	case a <- 1:
		// This is OK
	case <-b:
		// This is impossible
	case c <- 100:
		// This is impossible

		//default:
		//	a <- 10
	}
}

package main

import "fmt"

// <b>Global deadlock</b>: Program with (global) deadlock
//  ‣ 1 sender
//  ‣ 2 receivers (mismatch)

func main() {
	ch, done := make(chan int), make(chan int)
	go work()
	go send(ch)       // <b>Send</b> to ch.
	go recv(ch, done) // <b>Forward</b> from ch to done.
	go recv(ch, done) // <i>Who is ch receiving from?</i>
	print("Done:", <-done, <-done)
}

func send(ch chan int)       { ch <- 1 }
func recv(ch, done chan int) { done <- <-ch }

func work() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}

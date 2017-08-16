// +build ignore

package main

import (
	"fmt"
	"time"
)

func Gopher(speak chan<- string, listen <-chan string, done chan struct{}) {
	speak <- "Hey"
	fmt.Println("The other gopher said", <-listen)
	done <- struct{}{}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan struct{})
	// go makeCoffee()
	go Gopher(ch1, ch2, done) // Gopher 1 → ch1 → Gopher 2
	go Gopher(ch2, ch1, done) // Gopher 2 → ch2 → Gopher 1
	<-done
	<-done
}

func makeCoffee() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("coffee brewing...")
	}
}

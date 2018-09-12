package main

import "fmt"

func main() {
	ch := make(chan int)
	go looper()
	fmt.Println("Received", <-ch)
}

func looper() {
	for {
	}
}

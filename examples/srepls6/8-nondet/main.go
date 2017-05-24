package main

import "fmt"

func main() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("Received from case 1")
	case <-ch:
		fmt.Println("Received from case 2:")
	}
}

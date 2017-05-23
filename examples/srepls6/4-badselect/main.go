package main

import (
	"fmt"
)

func send(ch chan int, v int) {
	ch <- v
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1, 2)
	go send(ch2, 3)
	select {
	case x := <-ch1:
		fmt.Println("Case x=", x, "ch2=", <-ch2)
	case y := <-ch2:
		fmt.Println("Case y=", y) //,"ch1=", <-ch1)
	}
}

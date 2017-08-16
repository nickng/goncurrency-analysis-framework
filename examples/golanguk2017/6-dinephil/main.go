package main

// <b>The dining philosopher problem</b>
// By Edsger Dijkstra

import (
	"fmt"
)

func main() {
	fork1 := make(chan bool)
	fork2 := make(chan bool)
	fork3 := make(chan bool)
	go philo("Socrates", fork1, fork2)
	go philo("Plato", fork2, fork3)
	go philo("Aristotle", fork3, fork1)
	go aFork(fork1)
	go aFork(fork2)
	aFork(fork3)
}

func aFork(fork chan bool) {
	for {
		fork <- true
		<-fork
	}
}

// This is a philosopher, left/right are forks
func philo(name string, left, right chan bool) {
	for {
		select {
		case <-right: // Pick up right
			fmt.Println(name, "got right fork")
			select {
			case <-left:
				fmt.Println(name, "got left fork")
				left <- true
				right <- true
				fmt.Println(name, "Eats")
				//default:
				//	right <- true
			}
		case <-left: // Pick up left
			fmt.Println(name, "got left fork")
			select {
			case <-right:
				fmt.Println(name, "got right fork")
				left <- true
				right <- true
				fmt.Println(name, "Eats")
				//default:
				//	left <- true
			}
		}
	}
}

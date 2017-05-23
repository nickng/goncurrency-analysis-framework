package main

import "fmt"

func main() {
	fork1 := make(chan bool)
	fork2 := make(chan bool)
	fork3 := make(chan bool)
	go philo(fork1, fork2)
	go philo(fork3, fork1)
	go philo(fork2, fork3)
	go aFork(fork1)
	go aFork(fork2)
	go aFork(fork3)
	//	time.Sleep(10)
}

func aFork(fork chan bool) {
	for {
		fork <- true
		<-fork
	}
}

// This is a philosopher, left/right are forks
func philo(left, right chan bool) {
	for {
		select {
		case <-right: // Pick up right
			select {
			case <-left:
				left <- true
				right <- true
				fmt.Println("Eat")
			default:
				right <- true
			}
		case <-left: // Pick up left
			select {
			case <-right:
				left <- true
				right <- true
				fmt.Println("Eat")
			default:
				left <- true
			}
		}
	}
}

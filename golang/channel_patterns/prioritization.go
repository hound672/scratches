package main

import (
	"log"
	"time"
)

func producer(ch chan int) {
	time.Sleep(time.Second)
	for {
		ch <- 1
		time.Sleep(1 * time.Second)
	}
}

func main() {
	log.Println("Starting prioritization")

	ch1 := make(chan int) // higher priority
	ch2 := make(chan int)

	go producer(ch1)
	go producer(ch2)

	for {
		select {
		case <-ch1:
			log.Println("First select")
			return
		default:
		}

		select {
		case <-ch1:
			log.Println("Second select CH1")
			return
		case <-ch2:
			log.Println("Second select CH2")
			return
		}
	}
}

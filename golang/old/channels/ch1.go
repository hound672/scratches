package main

import (
	"log"
	"time"
)

func task(ch chan<- int ) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	log.Printf("Done")
}

func main() {
	log.Printf("start")

	ch := make(chan int)
	go task(ch)

	for i := range ch {
		log.Printf("i: %d", i)
	}
	log.Printf("done main")

	time.Sleep(2 * time.Second)
}

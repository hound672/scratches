package main

import (
	"log"
	"time"
)

func task() <- chan int {
	ch := make(chan int)

	go func() {
		log.Printf("before write")
		select {
		case ch <- 10:
			break
		default:
			break
		}
		log.Printf("after write")
	}()

	return ch
}

func main() {
	ch := task()

	_ = ch

	time.Sleep(time.Second * 5)
}

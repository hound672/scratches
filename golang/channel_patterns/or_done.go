package main

import (
	"log"
	"time"
)

func OrDone(done <-chan struct{}, in <- chan int) <-chan int{
	out := make(chan int)

	go func(){
		defer close(out)

		for {
		select {
		case <-done:
			return
		default:
		}

		select {
		case <-done:
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			out <- v
		}
		}
	}()

	return out
}

func main() {
	log.Println("Starting OrDone")

	in := make(chan int)

	go func (){
		for i := range 5 {
			time.Sleep(500*time.Millisecond)
			in <- i
		}
	}()

	done := make(chan struct{})
	go func (){
		time.Sleep(time.Second * 2)
		log.Println("close")
		close(done)
	}()

	for v := range OrDone(done, in) {
		log.Printf("data: %v", v)
	}
}

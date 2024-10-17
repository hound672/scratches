package main

import (
	"log"
	"sync"
)

func FanIN(channels ...<-chan int) <-chan int {
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	ch := make(chan int)
	
	for _, channel := range channels {
		go func() {
			defer wg.Done()
			for v := range channel {
				ch <- v
			}	
		}()
	}
	
	go func() {
		wg.Wait()
		close(ch)
	}()
	
	return ch
}

func main() {
	log.Printf("Starting Fan in")

	// prepare dummy channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		for i := 1; i <= 100; i++ {
			ch1 <- i
			ch2 <- i+1
			ch3 <- i+2
		}

	}()
	
	for v := range FanIN(ch1, ch2, ch3) {
		log.Println(v)
	}
}

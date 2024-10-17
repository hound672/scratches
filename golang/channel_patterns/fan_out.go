package main

import (
	"log"
	"sync"
	"time"
)

func FanOUT(inputCh <-chan int, count int) []chan int {
	if count <= 0 {
		count = 1
	}

	channels := make([]chan int, 0, count)
	for i := 0; i < count; i++ {
		ch := make(chan int)
		channels = append(channels, ch)
	}

	go func() {
		idx := 0
		for v := range inputCh {
			channels[idx] <- v
			idx = (idx + 1) % len(channels)
		}

		for _, ch := range channels {
			close(ch)
		}
	}()
		
	return channels
}

func main() {
	log.Printf("Starting fan out")

	ch := make(chan int)

	go func(){
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	channels := FanOUT(ch, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(){
		defer wg.Done()
		for v := range channels[0] {
			log.Printf("Channel 0 received %v", v)
		}
	}()
	go func(){
		defer wg.Done()
		for v := range channels[1] {
			log.Printf("Channel 1 received %v", v)
		}
	}()

	wg.Wait()
}

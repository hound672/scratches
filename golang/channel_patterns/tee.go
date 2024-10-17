package main

import (
	"log"
	"sync"
	"time"
)

func Tee(inputCh <-chan int, count int) []chan int {
	if count <= 0 {
		count = 1
	}
	
	channels := make([]chan int, 0, count)
	for i := 0; i < count; i++ {
		ch := make(chan int)
		channels = append(channels, ch)
	}

	go func() {
		defer func(){
			for _, ch := range channels {
				close(ch)
			}
		}()
		
		for v := range inputCh {
			for _, ch := range channels {
				ch<-v
			}
		}
	}()

	return channels
}

func main() {
	log.Printf("Starting tee")

	ch := make(chan int)

	go func(){
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	channels := Tee(ch, 2)
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

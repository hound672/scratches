package main

import (
	"log"
	"time"
)


func Or(channels ...<-chan struct{}) <-chan struct{}{
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan struct{})
	go func(){
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
			default:
			select {
			case <-orDone:
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-Or(channels[3:]...):
			}
		}
	}()
	
	return orDone
}

func main() {
	log.Printf("Starting or channels")

	after := func(after time.Duration) <- chan struct{} {
		c := make(chan struct{})
		go func(){
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	
	start := time.Now()

	<-Or(
		after(1*time.Second),
		after(2*time.Second),
		after(3*time.Second),
		after(4*time.Second),
		after(500*time.Millisecond),
		)

	log.Printf("Called after: %s", time.Since(start))
}

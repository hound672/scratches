package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	log.Printf("in go")
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Printf("done")
			return
		default:
			log.Printf("default")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func main() {
	log.Printf("start")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx, wg)

	time.Sleep(2 * time.Second)

	log.Printf("stop...")
	cancel()

	wg.Wait()
}

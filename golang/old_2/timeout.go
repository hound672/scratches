package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

func run(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	log.Printf("Start")

	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := run(ctx)
		log.Printf("err: %v", err)

		if errors.Is(err, context.Canceled) {
			log.Printf("canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("deadline")
		}

	}()

	_ = cancel
	// cancel()

	wg.Wait()
}

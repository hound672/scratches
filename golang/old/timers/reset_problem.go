package main

import (
	"fmt"
	"time"
)

func resetTimer(t *time.Timer, d time.Duration) {
	if !t.Stop() {
		select {
		case <-t.C:
		default:
		}
	}
	t.Reset(d)
}

func main() {
	const timeout = 10 * time.Millisecond
	t := time.NewTimer(timeout)
	time.Sleep(20 * time.Millisecond)

	start := time.Now()
	//t.Reset(timeout)
	resetTimer(t, timeout)
	<-t.C
	fmt.Printf("Time elapsed: %dms\n", time.Since(start).Milliseconds())
	// expected: Time elapsed: 10ms
	// actual:   Time elapsed:  0ms
}

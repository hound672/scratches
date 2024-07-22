package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Start\n")

	now := time.Now().UTC()
	next := now.Add(time.Second)

	fmt.Printf(" NOW: %v\n", now)
	fmt.Printf("NEXT: %v\n", next)

	delta := next.Sub(now)
	fmt.Printf("Delta: %v\n", delta < time.Second)
}

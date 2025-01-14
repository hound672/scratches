package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("Start time")

	t := time.Time{}
	log.Printf("Time: %v; isZero: %v", t, t.IsZero())

	now := time.Now()
	log.Printf("Time: %v; isZero: %v", now, now.IsZero())
}

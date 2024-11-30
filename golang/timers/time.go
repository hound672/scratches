package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	res := now.Format(time.RFC3339)
	log.Printf("res: [%s]", res)
}

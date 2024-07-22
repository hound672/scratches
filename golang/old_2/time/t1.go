package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("start")

	layout     := "2006-01-02T15:04:05.07"
	
	now := time.Now()
	
	res := now.Format(layout)
	log.Printf("res: %s", res)
	
}

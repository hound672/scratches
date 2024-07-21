package main

import (
	"log"
)

type Own struct {
	data []string
}

func main() {
	log.Printf("Start")

	o := &Own{}

	log.Printf("1 o: %v", o)
	
	o.data = append(o.data, []string{"1", "2", "3"}...)
	
	log.Printf("2 o: %v", o)
}

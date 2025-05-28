package main

import (
	"log"
)

func main() {
	log.Printf("start is nil")

	var v any
	v = "dsfsdf"

	if v == nil {
		log.Printf("v is nil")
	} else {
		log.Printf("v is not nil: %v", v)
	}
}

package main

import (
	"log"
)

func get(val *int) {
	if val == nil || *val == 123 {
		log.Println("GOT")
	}
}

func main() {
	log.Println("Start")

	v := 123
	_ = v
	get(&v)
}

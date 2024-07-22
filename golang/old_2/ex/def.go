package main

import (
	"log"
)

func f() {
	i := 0

	//defer log.Printf("i: %d", i)
	defer func() {
		log.Printf("i: %d", i)
	}()

	i = 5
}

func main() {
	log.Printf("Start")
	f()
}

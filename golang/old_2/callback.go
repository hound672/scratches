package main

import (
	"log"
)

type Base struct {
	f func()
}

func fff() {
	log.Printf("in func")
}

func main() {
	log.Printf("Start")

	b := &Base{
		//f: fff,
		//f: nil,
	}

	log.Printf("f: %v", b.f)
	if b.f != nil {
		b.f()
		log.Printf("Called")
		return
	}

	log.Printf("not func")
}

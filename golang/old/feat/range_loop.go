package main

import (
	"log"
)

func Backward[E any](s []E) func(func(E) bool) {
	return func(yield func(E) bool) {
		for i := len(s)-1; i >= 0; i-- {
			r := yield(s[i])
			log.Printf("r: %v", r)
			if !r {
				return
			}
		}
	}
}

func main() {
	log.Printf("Start")

	s := []string{"hello", "world"}
	for i := range Backward(s) {
		log.Println(i)
		//break
	}
}

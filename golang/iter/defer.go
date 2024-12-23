package main

import (
	"iter"
	"log"
)

func GetIter() iter.Seq[int] {
	defer func() {
		log.Printf("defer in root")
	}()

	return func(yield func(int) bool) {
		defer func() {
			log.Printf("Defer in returned func")
		}()
		for i := range 5 {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	log.Printf("Example with defer")
	
	for i := range GetIter() {
		log.Printf("Iter %d", i)
	}
}

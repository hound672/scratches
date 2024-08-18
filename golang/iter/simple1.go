package main

import (
	"iter"
	"log"
)

func SimpleIter() iter.Seq[int] {
	log.Println("Before")
	return func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
	}
}

func main() {
	for i := range SimpleIter() {
		log.Println(i)
	}
}

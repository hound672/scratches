package main

import (
	"iter"
	"log"
)

func SimpleIter() iter.Seq[int] {
	log.Println("Before")
	return func(yield func(int) bool) {
		defer func(){log.Println("Defer in yield func")}()
		yield(1)
		yield(2)
		yield(3)
	}
}

func main() {
	for i := range SimpleIter() {
		defer func() {log.Printf("defer: %v", i)}()
		log.Println(i)
	}
}

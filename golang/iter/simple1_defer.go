package main

import (
	"iter"
	"log"
)

func SimpleIter() iter.Seq[int] {
	defer log.Println("Defer in constructor")
	log.Println("Before")
	return func(yield func(int) bool) {
		defer func(){log.Println("Defer in yield func")}()
		yield(1)
		yield(2)
		yield(3)
	}
}

type Iter[E any] func(body func(index int, value E))

func Backward[S ~[]E, E any](s S) Iter[E] {
	return func(body func(int, E)) {
		for i := len(s) - 1; i >= 0; i-- {
			body(i, s[i])
		}
	}
}

func iterRange() {
	for i := range SimpleIter() {
		defer func() {log.Printf("defer: %v", i)}()
		log.Println("Before print i")
		log.Println(i)
		log.Println("After print i")
	}
	log.Println("Done in iterRange func")
}

func iterManual() {
	// Использование
	backwardIter := Backward([]int{1,2,3})

	backwardIter(func(index int, value int) {
		defer func(){log.Printf("defer: %v", value)}()

		log.Println("Before print i")
		log.Println(value)
		log.Println("After print i")
	})

	log.Println("Done in iterManual func")
}

func main() {
	iterRange()
	log.Println("------------------")
	iterManual()
}

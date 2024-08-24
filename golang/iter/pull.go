package main

import (
	"fmt"
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
	// push
	fmt.Printf("Push")
	for i := range SimpleIter() {
		fmt.Printf("i: %d\n", i)
	}

	// pull
	fmt.Println("Pull")
	simpleIterPull, stop := iter.Pull(SimpleIter())

	for i := range 5 {
		v, ok := simpleIterPull()
		if !ok {
			fmt.Println("not ok")
			return
		}
		fmt.Printf("i: %d; v: %d\n", i, v)

		if i == 1 {
			stop()
		}
	}
}
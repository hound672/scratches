package main

import (
	"fmt"
	"iter"
	"math/rand"
)

func main() {
	randIntGen := iter.Seq[int](func(yield func(int) bool) {
		for {
			if !yield(rand.Int()) {
				break
			}
		}
	})

	nextRandInt, stop := iter.Pull(randIntGen)
	defer stop()

	for i := range 5 {
		v, ok := nextRandInt()
		if !ok {
			fmt.Println("no more values in sequence")
			break
		}
		fmt.Printf("[%d] %d\n", i, v)
	}
}
package main

import (
	"iter"
	"log"
)

func SimpleIter() iter.Seq[int] {
	log.Println("Before")
	return func(yield func(int) bool) {
		for cnt := range 10 {
			if !yield(cnt) {
				log.Println("Cancled?")
				return
			}
			log.Println("After")
		}
	}
}

func main() {
	cnt := 0
	for i := range SimpleIter() {
		continue
		log.Println(i)
		cnt++
		if cnt >= 2 {
			break
		}
		continue
	}
}

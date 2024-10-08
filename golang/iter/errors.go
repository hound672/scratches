package main

import (
	"iter"
	"log"

	"github.com/go-faster/errors"
)

func getIter() iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		for cnt := range 10 {
			var e error
			if cnt > 3 {
				e = errors.New("test")
			}
			if !yield(cnt, e) {
				return
			}
		}
	}
}

func main() {
	log.Printf("Start")

	i := getIter()
	for v, err := range i {
		if err != nil {
			log.Printf("done")
			break
		}
		log.Printf("v: %v; err: %v", v, err)
	}
}

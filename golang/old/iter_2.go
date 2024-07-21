package main

import (
	"log"
	"time"
)

type S struct {
	Val int
}

func f(i *S) {
	log.Printf("in func: %d; val: %d", i, *i)
}
func main() {
	//sl := []int{1,2,3}
	sl := []S{
		{
		Val: 1,
		},
		{
			Val: 2,
		},
		{
			Val: 3,
		},
	}

	for _, el := range sl {
		log.Printf("i: %d", el)
		f(&el)
	}

	time.Sleep(2 * time.Second)
}

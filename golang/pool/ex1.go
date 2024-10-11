package main

import (
	"log"
	"sync"
)

type SomeData struct {
	ID int
}

var cnt int = 0

func NewSomeData() any {
	cnt++
	log.Printf("Const is called. cnt: %d", cnt)
	return &SomeData{ID: cnt}
}

func main() {
	log.Printf("Ex1 start")

	pool := &sync.Pool{
		New: NewSomeData,
	}

	obj1 := pool.Get()
	obj2 := pool.Get()

	log.Printf("obj1: %v; obj2: %v", obj1, obj2)

	pool.Put(obj1)
	obj3 := pool.Get()

	log.Printf("obj2: %v; obj3: %v", obj2, obj3)
}

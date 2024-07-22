package main

import (
	"log"
	"sync"
)

type OwnType struct {

}

func NewOwnType() interface{} {
	log.Printf("called")
	return &OwnType{}
}

var bufPool = sync.Pool{
	New: NewOwnType,
}

func main() {
	log.Printf("Start")

	v := bufPool.Get().(*OwnType)
	log.Printf("v: %v; %p", v, v)
	// bufPool.Put(v)

	v1 := bufPool.Get().(*OwnType)
	log.Printf("v: %v; %p", v1, v1)
}
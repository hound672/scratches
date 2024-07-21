package main

import (
	"log"
	"sync"
	"time"
)

var m sync.Mutex

func f(i int) {
	log.Printf("Before lock 1: ", i)

	m.Lock()

	log.Printf("Before lock 2: ", i)
	m.Unlock()

	m.Lock()

	log.Printf("After lock 2", i)
	m.Unlock()
}

func main() {
	log.Printf("Start ")

	go f(1)
	go f(2)

	time.Sleep(time.Second * 5)
}

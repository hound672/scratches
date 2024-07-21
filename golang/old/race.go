package main

import (
	"log"
	"sync"
)


type Settings struct {
	Val int
	S string
}

func main() {
	log.Printf("start")

	m := &Settings{}
	mut := sync.RWMutex{}
	_ = mut

	write := func() {
		mut.Lock()
		defer mut.Unlock()
		m.Val = 11
		m.S = "sdfsdf"
	}

	read := func() {
		mut.RLock()
		defer mut.RUnlock()
		_ = m.S
		_ = m.Val
	}

	for i := 0; i < 100; i ++ {
		go read()
		go write()
	}
}

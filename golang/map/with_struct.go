package main

import (
	"log"
)

type Record struct {
	Old string
	New string
}

func main() {
	log.Printf("Start")
	
	m := make(map[string]Record)

	rec, ok := m["NEW"]
	if !ok {
		m["NEW"] = Record{}
		rec = m["NEW"]
	}
	rec.Old = "OLD"
	//m["NEW"] = rec

	log.Printf("map: %v", m)
}

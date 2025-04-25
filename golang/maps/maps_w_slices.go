package main

import (
	"log"
)

func main() {
	log.Printf("Start")
	
	m := map[int][]int{}

	log.Printf("before: %v", m)

	m[1] = append(m[1], 1)
	m[1] = append(m[1], 2)

	log.Printf("after: %v", m)
}

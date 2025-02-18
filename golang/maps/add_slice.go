package main

import (
	"log"
)

func main() {
	log.Printf("Start")

	m := map[int][]string{}

	m[1] = append(m[1], "111")
	m[1] = append(m[1], "111")

	m[2] = append(m[2], "222")

	log.Printf("m: %v", m)
}

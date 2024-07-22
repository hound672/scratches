package main

import "log"

func f(members ...interface{}) {
	log.Printf("mem: %v; len(%d)", members, len(members))
}

func main() {
	log.Printf("Hello")

	d := []string{"one", "two", "three"}
	r := make([]interface{}, 0, len(d))

	for i := 0; i < len(d); i++ {
		r = append(r, d[i])
	}

	f(d)
	f(r...)
}

package main

import "log"

func main() {
	log.Printf("Start")

	var i int8 = 120
	i += 10
	log.Printf("i: %v", i)
}

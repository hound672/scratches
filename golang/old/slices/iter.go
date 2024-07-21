package main

import "log"

func main() {
	var sl []int
	//sl = append(sl, 1)
	//sl = append(sl, 2)
	//sl = append(sl, 3)

	log.Printf("sl: %v; len: %d; cap: %d", sl, len(sl), cap(sl))

	for _, el := range sl {
		log.Printf("el: %v", el)
	}
}

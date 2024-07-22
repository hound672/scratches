package main

import "log"

func main() {
	log.Printf("Start")

	ar := [3]int{0, 1, 2}
	log.Printf("ar: %T; %v, len: %d; cap: %d", ar, ar, len(ar), cap(ar))

	sl := ar[:]
	log.Printf("ar: %T; %v, len: %d; cap: %d", ar, ar, len(sl), cap(sl))


	s := make([]int, 0, 5)
	log.Printf("t: %T", s)
}

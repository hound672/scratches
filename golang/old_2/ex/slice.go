package main

import "log"

func f(sl []int) {
	//log.Printf("--------------")
	//log.Printf("addr: %p", sl)
	sl = append(sl, 1,2,3)
	sl[0] = 222
	log.Printf("s: %v", sl)
	//log.Printf("addr: %p", sl)
	//log.Printf("--------------")
}

func main() {
	log.Printf("Start")

	sl := make([]int, 1, 5)
	sl[0] = 111

	log.Printf("addr: %p", sl)
	log.Printf("cont: %v", sl)

	f(sl)

	log.Printf("addr: %p", sl)
	log.Printf("cont: %v", sl)

	//sl1 := sl[:]
	//log.Printf("addr: %p", sl1)
}

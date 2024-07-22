package main

import "log"

func main() {
	log.Printf("slice start")

	arr := [2]int{1,2}


	log.Printf("%v; %T; len: %d, cap: %d; %p", arr, arr, len(arr), cap(arr), &arr[0])

	sl := arr[:]
	log.Printf("%v; %T; len: %d, cap: %d; %p", sl, sl, len(sl), cap(sl), sl)

	log.Printf("--------------")

	sl[0] = 111

	log.Printf("%v; %T; len: %d, cap: %d", arr, arr, len(arr), cap(arr))
	log.Printf("%v; %T; len: %d, cap: %d", sl, sl, len(sl), cap(sl))

	log.Printf("--------------")

	sl = append(sl, 3)

	log.Printf("%v; %T; len: %d, cap: %d", arr, arr, len(arr), cap(arr))
	log.Printf("%v; %T; len: %d, cap: %d", sl, sl, len(sl), cap(sl))

	log.Printf("--------------")

	sl[1] = 222

	log.Printf("%v; %T; len: %d, cap: %d; %p", arr, arr, len(arr), cap(arr), &arr[0])
	log.Printf("%v; %T; len: %d, cap: %d; %p", sl, sl, len(sl), cap(sl), sl)
}

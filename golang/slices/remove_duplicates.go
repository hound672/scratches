package main

import (
	"log"
)

func RemoveDuplicates[T comparable](strSlice []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	sl := []int{1,2,3,4, 1}
	log.Printf("before: %v", sl)
	res := RemoveDuplicates(sl)
	log.Printf("after: %v; res: %v", sl, res)

	slS := []string{"one", "two", "three", "one"}
	log.Printf("sl: %v", slS)
	resS := RemoveDuplicates(slS)
	log.Printf("after: %v; res: %v", slS, resS)
}
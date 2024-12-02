package main

import (
	"log"
	"reflect"
	"slices"
)

func main() {
	log.Println("Compare")

	sl1 := []string{"one", "two", "three"}
	sl2 := []string{"two", "three", "one"}

	log.Printf("Before 1: %v", slices.Equal(sl1, sl2))
	log.Printf("Before 2: %v", reflect.DeepEqual(sl1, sl2))

	slices.Sort(sl1)
	slices.Sort(sl2)

	res := slices.Equal(sl1, sl2)
	log.Println(res)
}


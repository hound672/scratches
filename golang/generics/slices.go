package main

import (
	"log"
	"slices"
)

func IsSlicesEquals(a, b []any) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	log.Printf("Start generics.slices")

	sl1 := []string{"111", "222"}
	sl2 := []string{"111", "222"}

	res := slices.Equal(sl1, sl2)
	log.Printf("res: %v", res)
}

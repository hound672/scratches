package main

import (
	"log"

	"github.com/google/uuid"
)

func interSection[T comparable](pS ...[]T) []T {
	hash := make(map[T]*int) // value, counter
	result := make([]T, 0)
	for _, slice := range pS {
		duplicationHash := make(map[T]bool) // duplication checking for individual slice
		for _, value := range slice {
			if _, isDup := duplicationHash[value]; !isDup { // is not duplicated in slice
				if counter := hash[value]; counter != nil { // is found in hash counter map
					if *counter++; *counter >= len(pS) { // is found in every slice
						result = append(result, value)
					}
				} else { // not found in hash counter map
					i := 1
					hash[value] = &i
				}
				duplicationHash[value] = true
			}
		}
	}
	return result
}

func difference[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []T
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func main() {
	log.Println("Start")

	sl1 := []string{"123", "456", "789"}
	sl2 := []string{"123", "456"}
	log.Printf("Before: sl1: %v; sl2: %v", sl1, sl2)

	res := interSection(sl1, sl2)
	log.Printf("Itersection: %v", res)

	log.Printf("After: sl1: %v; sl2: %v", sl1, sl2)

	log.Println("---------------------")

	resultSlice := []string{}
	checkMap := map[string]struct{}{}

	for _, addr := range sl1 {
		checkMap[addr] = struct{}{}
	}
	for _, addr := range sl2 {
		if _, ok := checkMap[addr]; ok {
			resultSlice = append(resultSlice, addr)
		}
	}

	log.Println(resultSlice)

	log.Println("---------------------")

	log.Printf("Before: sl1: %v; sl2: %v", sl1, sl2)

	r := difference(sl1, sl2)
	log.Println(r)

	log.Printf("After: sl1: %v; sl2: %v", sl1, sl2)

	log.Println("---------------------")
	u1 := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	u2 := []uuid.UUID{u1[0], u1[1], uuid.New()}
	log.Printf("Before: u1: %v; u2: %v", u1, u2)
	dif := difference(u1, u2)
	log.Println(dif)
	log.Printf("After: u1: %v; u2: %v", u1, u2)

	log.Println("---------------------")
	i := interSection(u1, u2)
	log.Println(i)
}

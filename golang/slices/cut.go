package main

import (
	"log"
	"strings"
)

func main() {
	log.Printf("Start slice.cut")
	
	source := []string{"1","2","3","4","5","6","7","8","9"}
	res1 := strings.Join(source[0:3], "\n")
	res2 := source[3:5]
	
	log.Printf("res1: %v, res2: %v", res1, res2)
}

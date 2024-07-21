package main

import (
	"log"
	"strconv"
)

func main() {
	log.Printf("Start")
	
	source := true
	res := strconv.FormatBool(int(source))
	
	log.Printf("res: %v", res)
}

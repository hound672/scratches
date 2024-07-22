package main

import (
	"log"
	"strings"
)

func main() {
	log.Printf("Start")

	source := "mail"

	res := strings.SplitN(source, "@", 2)
	log.Printf("res: %v", res)
}

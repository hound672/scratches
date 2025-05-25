package main

import (
	"crypto/sha256"
	"log"
)

func main() {
	log.Printf("Start sha256")

	name := "foobar"
	res := sha256.Sum256([]byte(name))

	log.Printf("SHA256: %x", res)
}


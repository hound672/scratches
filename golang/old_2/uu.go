package main

import (
	"log"

	"github.com/google/uuid"
)

func main() {
	log.Printf("Start")

	u, _ := uuid.NewUUID()

	log.Printf("UUID: %s", u)
}

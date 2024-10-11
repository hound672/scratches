package main

import (
	"log"

	"github.com/google/uuid"
)

func main() {
	log.Printf("Start")

	newUUID, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	log.Printf("UUID: %s", newUUID.String())

	var id uuid.UUID
	id, err = uuid.Parse(newUUID.String())
	if err != nil {
		panic(err)
	}

	log.Printf("Filled: %s", id)
}

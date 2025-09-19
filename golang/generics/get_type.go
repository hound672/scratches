package main

import (
	"log"
)

func getType[T any](source string) T {
	switch T.(type) {

	}
	return *new(T)
}

func main() {
	log.Printf("Start get type")
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Reader()string {
	// read from smth
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	response := fmt.Sprintf("Response #: %d", rand.Int())
	return response
}

func GetResult(addresses []string) string {
	result := make(chan string)

	for _, address := range addresses {
		go func(address string) {
			response := Reader()
			select {
			case result <- response:
				log.Printf("Response from address: %s", address)
			default:
				return
			}
		}(address)
	}

	return <-result
}

func main() {
	log.Printf("Starting moving later")

	addresses := []string{"add_1", "add_2", "add_3"}

	value := GetResult(addresses)
	log.Printf("Result: %s", value)
}

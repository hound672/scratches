package main

import (
	"errors"
	"log"
)

func rErr1() error {
	return errors.New("some error 1")
}

func rErr2() error {
	return errors.New("some error 2")
}


func main() {
	log.Printf("Start if")

	err := rErr1()
	log.Printf("err: %v", err)

	if err = rErr2(); err != nil {
		log.Printf("err: %v", err)
	}

	log.Printf("again err: %v", err)
}


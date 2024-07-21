package main

import (
	"errors"
	"fmt"
	"log"
)

var Err1 = errors.New("error_1")
var Err2 = errors.New("error_2")

func f() error {
	return fmt.Errorf("some error: %w; %w", Err1, Err2)
}

func f2() error {
	e1 := fmt.Errorf("err1: %w", Err1)
	e2 := fmt.Errorf("err2: %w", e1)
	return e2
}

func main() {
	err := fmt.Errorf("%w %w", Err1, Err2)
	log.Printf("errors.Is(err, Err1): %v", errors.Is(err, Err1)) // true
	log.Printf("errors.Is(err, Err2): %v", errors.Is(err, Err2)) // true

	log.Printf("Start")

	err = f()
	log.Printf("err: %v", err)

	if errors.Is(err, Err1) {
		log.Printf("is err1")
	}
	if errors.Is(err, Err2) {
		log.Printf("is err2")
	}
	
	log.Printf("======================")
	
	err = f2()
	log.Printf("err2: %v", err)
	
	if errors.Is(err, Err1) {
		log.Printf("is err1")
	}
}

package main

import (
	"errors"
	"fmt"
	"log"
)

var Err1 = errors.New("one")
var Err2 = errors.New("two")
var Err3 = errors.New("three")

func handler(key string) error {
	errMap := map[string]error{
		"one": Err1,
		"two": Err2,
		"three": Err3,
	}

	e, ok := errMap[key]
	if !ok {
		return errors.New("unknown")
	}

	return fmt.Errorf("%w: some error", e)
}

func main() {
	log.Printf("Start")

 	err := handler("three1")
	 log.Printf("res: %v", err)

	 switch {
	 case err == nil:
		 log.Printf("all fine")
	 case errors.Is(err, Err1):
		 log.Printf("ERR1")
	 case errors.Is(err, Err2):
		 log.Printf("ERR2")
	 case errors.Is(err, Err3):
		 log.Printf("ERR3")

	 }
}
package main

import (
	"errors"
	"log"
)

var (
	Err1 = errors.New("error 1")
	Err2 = errors.New("error 2")
	Err3 = errors.New("error 3")
)

func main() {
	err := Err1

	switch {
	case errors.Is(err, Err1):
	case errors.Is(err, Err2):
		log.Printf("is err1 or err2")
	case errors.Is(err, Err3):
		log.Printf("is err3 or Err3")
	}
}

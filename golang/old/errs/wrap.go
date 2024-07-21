package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrMain1 = errors.New("main 1")
var ErrMain2 = errors.New("main 2")
var ErrMain3 = errors.New("main 3")

func f1() error {
	return ErrMain1
}

func f2() error {
	err := f1()
	return fmt.Errorf("%w:%w", ErrMain2, err)
}

func main() {
	log.Printf("Start")
	
	e1 := f1()
	e2 := f2()
	
	if errors.Is(e1, ErrMain1) {
		log.Printf("errors.Is(e1, ErrMain1)")
	}
	if errors.Is(e2, ErrMain1) {
		log.Printf("errors.Is(e2, ErrMain1)")
	}
	if errors.Is(e2, ErrMain2) {
		log.Printf("errors.Is(e2, ErrMain2)")
	}

}

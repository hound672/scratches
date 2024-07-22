package main

import (
	"errors"
	"fmt"
)

var Nerr = errors.New("error own")

func own() error {
	return Nerr
	// return nil
	// return errors.New("dsfsdf")
}

func main() {
	fmt.Printf("start\n")

	err := own()

	if !errors.Is(err, Nerr) {
		fmt.Printf("is err!")
	}
}

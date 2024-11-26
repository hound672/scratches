package main

import (
	"errors"
	"log"

	"fmt"
)

type SomeError struct {
	Val string
}

func (e *SomeError) Error() string {
	return fmt.Sprintf("some error: %v", e.Val)
}

func rErr() error {
	return &SomeError{Val: "rErr"}
}

func rErr2() error {
	err := rErr()
	if err != nil {
		return fmt.Errorf("[WRAPPED] rErr2: %w", err)
	}

	return nil
}

func main() {
	err := rErr2()

	if err != nil {
		log.Println(err)

		var e *SomeError
		if errors.As(err, &e) {
			log.Println(e.Val)
		}
	}
}

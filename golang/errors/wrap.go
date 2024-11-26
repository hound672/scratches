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

var ErrSmth = errors.New("something went wrong")

func rErr() error {
	return &SomeError{Val: "rErr"}
	//return ErrSmth
}

func rErr2() error {
	err := rErr()
	if err != nil {
		return fmt.Errorf("[WRAPPED] rErr2: %w", err)
	}

	return nil
}

func rErr3() error {
	err := rErr2()
	if err != nil {
		return fmt.Errorf("[WRAPPED] rErr3: %w", err)
	}

	return nil
}

func main() {
	err := rErr3()

	if err != nil {
		log.Println(err)

		var e *SomeError

		switch {
		case errors.Is(err, ErrSmth):
			log.Printf("ErrSmth: %v", err)
		case errors.As(err, &e):
			log.Printf("ErrAs: %v; val: %v", e, e.Val)
		}
	}
}

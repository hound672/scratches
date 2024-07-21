package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	UUID string
	Username string
}

func main() {
	user := &User{}
	gofakeit.Struct(user)
	fmt.Printf("user: %v\n", user)

	fmt.Printf("beer: %s", gofakeit.BeerName())

	s := gofakeit.Word()
	fmt.Printf("s: %s", s)

	s1 := gofakeit.Word()
	fmt.Printf("s: %s", s1)

	gender := gofakeit.Gender()
	fmt.Printf("gender: %s", gender)
}
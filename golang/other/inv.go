package main

import (
	"log"
)

type Inter interface {
	GetValue() int
}

type SomeType struct {
	Value int
}

func (s *SomeType) GetValue() int {
	return s.Value
}

func get() []Inter {
	return []Inter{&SomeType{}}
}

func accept(i []Inter) {}

func main() {
	obj := []*SomeType{&SomeType{}}
	
	sl := []Inter{&SomeType{}}

	accept(sl)
	accept(obj)
	
	log.Printf("Start")
}

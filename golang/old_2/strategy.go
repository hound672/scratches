package main

import "fmt"

type Strategy interface {
	Get()
}

func NewStrategy() Strategy {
	strategy := &cloudStrategy{}
	return strategy
}

type commonStrategy struct {
	
}

func (cs *commonStrategy) Get() {
	fmt.Printf("called common strategy\n")
}

type cloudStrategy struct {
	commonStrategy
}

func (cs *cloudStrategy) Get() {
	fmt.Printf("called cloud strategy\n")
}

func main() {
	fmt.Printf("Start\n")
	
	strategy := NewStrategy()
	
	strategy.Get()
}

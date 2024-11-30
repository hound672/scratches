package main

import (
	"log"
)

type base struct {
	Val string
}

func (b *base) Process() {
	log.Printf("base.Val: %v", b.Val)
}

type Child struct {
	base
}

func main() {
	log.Printf("Start overide")

	child := &Child{base{
		Val: "some",
	}}

	child.Process()
}

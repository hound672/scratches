package main

import (
	"log"
)

type Cls1 struct {

}

func (c *Cls1) Foo() {
	log.Printf("In foo")
}

type Cls2 struct {
	Cls1
}

// func (c *Cls2) Foo() {
// 	log.Printf("In foo cls2")
// }

func main() {
	log.Printf("Start")

	// c := &Cls1{}
	c := &Cls2{}

	c.Foo()
}
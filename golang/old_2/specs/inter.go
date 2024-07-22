package main

import (
	"log"
)

type BaseInterface interface {
	Foo()
}

type BaseImpl struct {
	BaseInterface
}

func (b *BaseImpl) Foo() {
	log.Printf("called foo")
}

func NewBaseInterface() BaseInterface {
	impl := &BaseImpl{}
	return impl
}

func main() {
	log.Printf("start")

	impl := NewBaseInterface()
	impl.Foo()

	impl2 := &BaseImpl{}
	impl2.Foo()
}

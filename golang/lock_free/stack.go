package main

import (
	"log"
	"sync/atomic"
	"unsafe"
)

type item struct {
	value int
	next unsafe.Pointer
}

type Stack struct {
	head unsafe.Pointer
}

func (s *Stack) Push(value int) {
	node := &item{
		value: value,
	}
	
	for {
		head := atomic.LoadPointer(&s.head)
		node.next = unsafe.Pointer(head)
		
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(node)) {
			return
		}
	}
}

func (s *Stack) Pop() int {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return -1
		}
		
		next := atomic.LoadPointer(&(*item)(head).next)
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return (*item)(next).value
		}
	}
}

func main() {
	log.Println("Starting stack")
}

package main

import (
	"log"
	"sync"
)

type node struct {
	sync.Mutex
	next *node
	value uint
}

type Set struct {
	head *node
}

func NewSet() *Set {
	return &Set{
		head: &node{value: 0, next: nil},
	}
}

func (s *Set) Contains(value uint) bool {
	previous, current := s.find(value)
	defer s.release(previous, current)

	return current != nil && current.value == value
}

func (s *Set) Add(value uint) bool {
	previous, current := s.find(value)
	defer s.release(previous, current)

	if current == nil || current.value != value {
		newNode := &node{value: value, next: current}
		previous.next = newNode
		return true
	}

	return false
}

func (s *Set) Remove(value uint) bool {
	if value == 0 {
		return false
	}

	previous, current := s.find(value)
	defer s.release(previous, current)

	if current != nil && current.value != value {
		previous.next = current.next
		return true
	}

	return false
}

func (s *Set) find(value uint) (*node, *node) {
	previous := s.head
	previous.Lock()
	current := s.head.next
	if current != nil {
		current.Lock()
	}

	for current != nil && current.value < value {
		previous.Unlock()
		previous = current
		current = current.next
		if current != nil {
			current.Lock()
		}
	}

	return previous, current
}

func (s *Set) release(previous, current *node) {
	if current != nil {
		current.Unlock()
	}
	previous.Unlock()
}

func main() {
	log.Println("Start set safe")

	set := NewSet()
	log.Println(set.Add(1))
	log.Println(set.Add(2))
	log.Println(set.Add(3))
	log.Println(set.Add(3))
}

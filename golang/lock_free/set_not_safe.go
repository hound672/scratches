package main

import (
	"log"
)

type node struct {
	next *node
	value uint
}

type Set struct {
	head *node
}

func NewSet() *Set {
	return &Set{
		head: &node{value: 0},
	}
}

func (s *Set) Contains(value uint) bool {
	_, current := s.find(value)
	return current != nil && current.value == value
}

func (s *Set) Add(value uint) bool {
	if value == 0 {
		return false
	}

	previous, current := s.find(value)
	if current != nil && current.value == value {
		return false
	}

	newNode := &node{value: value, next: current}
	previous.next = newNode
	return true
}

func (s *Set) Remove(value uint) bool {
	if value == 0 {
		return false
	}

	previous, current := s.find(value)
	if current == nil || current.value != value {
		return false
	}

	previous.next = current.next
	return true
}

func (s *Set) find(value uint) (*node, *node) {
	previous := s.head
	current := s.head.next

	if current != nil && current.value < value {
		previous = current
		current = current.next
	}

	return previous, current
}

func main() {
	log.Println("Starting set not safe")

	set := NewSet()
	log.Println(set.Add(1))
	log.Println(set.Add(2))
	log.Println(set.Add(3))
	log.Println(set.Add(3))
}

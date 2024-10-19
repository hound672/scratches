package main

import (
	"log"
)

type itemNotSafe struct {
	value int
	next *itemNotSafe
}

type QueueNotSafe struct {
	head *itemNotSafe
	tail *itemNotSafe
}

func NewQueueNotSafe() *QueueNotSafe {
	dummy := &itemNotSafe{}
	return &QueueNotSafe{
		head: dummy,
		tail: dummy,
	}
}

func (q *QueueNotSafe) Push(v int) {
	q.tail.next = &itemNotSafe{value: v}
	q.tail = q.tail.next
}

func (q *QueueNotSafe) Pop() int {
	if q.head == q.tail {
		return -1
	}

	value := q.head.next.value
	q.head = q.head.next
	return value
}

func main() {
	queue := NewQueueNotSafe()
	queue.Push(1)
	queue.Push(2)
	v := queue.Pop()
	log.Printf("Pop: %v", v)
	v = queue.Pop()
	log.Printf("Pop: %v", v)
}
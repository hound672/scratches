// semaphore implementation
package main

import (
	"log"
	"sync"
	"time"
)

type Semaphore struct {
	count int
	max int
	cond *sync.Cond
}

func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		count: 0,
		max: max,
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Semaphore) Acquire() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	if s.count >= s.max {
		s.cond.Wait()
	}
	s.count++
}

func (s *Semaphore) Release() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	
	s.count--
	s.cond.Signal()
}

func work(number int, s *Semaphore) {
	log.Printf("worker #%d; BEFORE time: %v", number, time.Now())
	
	s.Acquire()

	log.Printf("worker #%d; AFTER time: %v", number, time.Now())
	time.Sleep(2 * time.Second)
	log.Printf("worker #%d; END time: %v", number, time.Now())
	s.Release()
}

func main() {
	log.Printf("Ex2")
	
	s := NewSemaphore(3)
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		work(1, s)
	}()
	go func() {
		defer wg.Done()
		work(2, s)
	}()
	go func() {
		defer wg.Done()
		work(3, s)
	}()
	go func() {
		defer wg.Done()
		work(4, s)
	}()

	wg.Wait()
}

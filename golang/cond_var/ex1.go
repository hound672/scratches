package main

import (
	"log"
	"sync"
	"time"
)

var data map[string]struct{}

func subscribe(name string, c *sync.Cond) {
	c.L.Lock()

	for len(data) == 0 {
		c.Wait()
	}

	log.Printf("Subscribing to %s", name)
	c.L.Unlock()
}

func publish(c *sync.Cond) {
	time.Sleep(time.Second * 2)
	log.Printf("Change")

	c.L.Lock()
	data["value"] = struct{}{}
	c.L.Unlock()

	log.Printf("Send broadcast")
	c.Broadcast()
}

func main() {
	log.Printf("Start")

	data = make(map[string]struct{})
	cond := sync.NewCond(&sync.Mutex{})

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		subscribe("subs 1", cond)
	}()

	go func() {
		defer wg.Done()
		subscribe("subs 2", cond)
	}()

	go func() {
		defer wg.Done()
		publish(cond)
	}()

	wg.Wait()
}

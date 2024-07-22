package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("Start")

	cnt := 0

	log.Println("before loop")
loop:
	for {
		log.Println("beggin in loop")
		cnt += 1

		switch {
		case cnt%2 == 0:
			log.Printf("cnt is even")
		case cnt > 5:
			log.Printf("cnt is more than 5")
			break loop
		}
		time.Sleep(time.Second)

	}

	log.Println("after loop")
}

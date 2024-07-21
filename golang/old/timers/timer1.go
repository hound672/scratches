package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("Start")

	//timer := time.NewTimer(time.Second* 2)
	//
	//for {
	//	select {
	//	case <- timer.C:
	//		log.Printf("timer")
	//		timer.Reset(time.Second * 2)
	//	default:
	//	}
	//}

	ticker := time.NewTicker(time.Second * 2)

	for {
		select {
		case <- ticker.C:
			log.Printf("Tick...")
		}
	}
}

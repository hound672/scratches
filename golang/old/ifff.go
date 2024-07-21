package main

import "log"

func check() bool {
	log.Printf("CHECK is called")
	return true
}

func main() {
	log.Printf("!@!!")

	cond := false

	if cond && check() {

	}
}

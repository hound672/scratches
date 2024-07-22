package main

import (
	"fmt"
)

func end(val int) {
	fmt.Printf("end: %d", val)
}

func run() {
	val := 10

	defer func(){
		fmt.Printf("before end: %v\n", val)
		defer end(val)
	}()

	val = 11
}

func main() {
	fmt.Printf("start\n")
	run()
}

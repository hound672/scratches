package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("i: %d\n", i)
		}(i)
		runtime.Gosched()
	}
}

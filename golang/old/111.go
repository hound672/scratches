package main

import "fmt"

func main() {
	num := make(chan int)

	values := []int{1, 2, 3}
	for _, v := range values {
		go func() {
			num <- v
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		fmt.Println(<-num)
	}
}

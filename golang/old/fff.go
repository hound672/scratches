package main

import (
	"fmt"
	"time"
)

func r(val string) {
	fmt.Printf("val: %v\n", val)
}

func main() {
	m := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	
	for _, v := range m {
		go r(v)
	}
	
	time.Sleep(1 * time.Second)
}

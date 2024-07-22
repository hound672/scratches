package main

import (
	"fmt"
	"reflect"
)

type Str struct {
	Field *bool `json:"field,omitempty"`
}

func main() {
	//map1 := make(map[string]map[int]int)
	map2 := make(map[string]map[int]int)

	map1 := map[string]map[int]int{
		"one": {1: 2},
	}

	//map1["one"] = make(map[int]int)
	//map1["one"][1] = 2

	map2["one"] = make(map[int]int)
	map2["one"][1] = 2

	fmt.Printf("is: %v\n", reflect.DeepEqual(map1, map2))
}

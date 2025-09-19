package main

import (
	"encoding/json"
	"log"
)

type SomeStruct struct {
	Values map[string]int
}

func main() {
	log.Printf("map marshall")
	
	values := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	someStruct := SomeStruct{Values: values}
	
	result, err := json.Marshal(someStruct)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("struct: %v", someStruct)
	log.Printf("marshalled: %v", string(result))
	
	unmarshalled := &SomeStruct{}
	if err := json.Unmarshal(result, &unmarshalled); err != nil {
		log.Fatal(err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)
	log.Printf("unmarshalled: %v", unmarshalled.Values["one"])
}

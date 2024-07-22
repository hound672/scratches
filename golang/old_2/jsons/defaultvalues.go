package main

import (
	"encoding/json"
	"log"
)

type PaginationParams struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}

func main() {
	log.Printf("Start")

	rawData := []byte(`{"limit": 123, "offset": "456"}`)

	s := &PaginationParams{Offset: 0, Limit: 0}
	err := json.Unmarshal(rawData, s)
	if err != nil {
		log.Printf("err: %v", err)
	}

	log.Printf("res: %v", s)
}

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("hello")
	
	cl := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:50002", http.NoBody)
	if err != nil {
		log.Fatalf("new: %v", err)
	}
	
	_, err = cl.Do(req)
	if err != nil {
		log.Fatalf("send: %s", err)
	}
}

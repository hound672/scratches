package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request; headers: %v", r.Header)
}

func main() {
	log.Printf("Start server")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

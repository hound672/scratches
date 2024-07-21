package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Printf("Start")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:1111", http.NoBody)
	if err != nil {
		log.Fatalf("err new req: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("err do: %v", err)
	}
	
	log.Printf("resp: %v", resp)

}
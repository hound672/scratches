package main

import (
	"context"
	"log"
	"time"

	"net/http"
	_ "net/http/pprof"
)

func smth() {
	for {
		time.Sleep(time.Second)
	}
}

func cont() {
	ctx, _ := context.WithTimeout(context.Background(), time.Hour)

	_ = ctx
	time.Sleep(time.Hour)
}

func main() {
	log.Printf("Start")

	go smth()
	go cont()

	_ = http.ListenAndServe("localhost:8080", nil)
}

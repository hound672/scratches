package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var cnt int

func main() {
	cnt = 0
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	cnt += 1
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

	delay := rand.Intn(1000)

	log.Printf("client: %v; delay: %v", cnt, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
}

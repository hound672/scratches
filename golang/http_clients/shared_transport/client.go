package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

var transport *http.Transport
var client1 *http.Client

func init() {
	transport = http.DefaultTransport.(*http.Transport).Clone()
}

func send() {
	log.Printf("Sending request to server")

	req1, err := http.NewRequest("GET", "http://localhost:8080/", bytes.NewBuffer(nil))
	if err != nil {
		panic(err)
	}

	_, err = client1.Do(req1)
	if err != nil {
		panic(err)
	}	
}

func main() {
	log.Printf("Run client")
	client1 = &http.Client{}

	for range 10 {
		send()
		time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Hour)

	//client1 := http.Client{Transport: transport}
	client1 := http.Client{}
	//client2 := http.Client{Transport: transport}
	client2 := http.Client{}

	req1, err := http.NewRequest("GET", "http://localhost:8080/", bytes.NewBuffer(nil))
	if err != nil {
		panic(err)
	}
	req1.Header.Set("Auth", "Hello1")
	_, err = client1.Do(req1)
	if err != nil {
		panic(err)
	}

	req2, err := http.NewRequest("GET", "http://localhost:8080/", bytes.NewBuffer(nil))
	if err != nil {
		panic(err)
	}
	req2.Header.Set("Auth", "Hello2")
	_, err = client2.Do(req2)
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Hour)
}

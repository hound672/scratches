package main

import (
	"bytes"
	"log"
	"net/http"
)

var transport *http.Transport

func init() {
	transport = http.DefaultTransport.(*http.Transport).Clone()
}

func main() {
	log.Printf("Run client")

	client1 := http.Client{Transport: transport}
	client2 := http.Client{Transport: transport}

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

}

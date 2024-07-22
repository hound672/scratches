package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Printf("Start")
	
	res, err := os.MkdirTemp("/tmp", "")
	if err != nil {
		log.Fatalf("err create: %v", err)
	}
	
	log.Printf("res: %v", res)
	
	time.Sleep(time.Second * 10)
	
	log.Printf("delete...")
	
	err = os.RemoveAll(res)
	if err != nil {
		log.Fatalf("error remove: %v", err)
	}
}

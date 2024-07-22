package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Printf("start")

	log.Printf("Start open file")
	f, err := os.OpenFile("/tmp/tmp.rntBMHSOv3/file.lock", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("OpenFile: %v", err)

	}
	log.Printf("Opened")

	time.Sleep(time.Second * 10)
	if err := f.Close(); err != nil {
		log.Fatalf("f.Close: %v", err)
	}
	log.Printf("Done")
}

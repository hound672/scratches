package main

import (
	"log"
	"os"
)

func main() {
	err := os.Rename("/tmp/tmp.xxKtDzCC8a/main.txt", "/tmp/tmp.PqPYoxmoZb/main.txt")
	if err != nil {
		log.Printf("Err: %v", err)
		return
	}
	log.Printf("all fine")
}

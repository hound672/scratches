package main

import (
	"log"
	"os"

	"golang.org/x/mod/sumdb/dirhash"
)

func main() {
	log.Printf("Hello")

	// path := "/tmp/easyp/mod/github.com/grpc-ecosystem/grpc-gateway/0bcc6bf00e0bf6ac71caab131df54e13af377603"
	path := "/tmp/easyp"

	res, err := dirhash.HashDir(path, "", dirhash.DefaultHash)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("NOT EXISTS")
		}
		log.Fatal(err)
	}

	log.Printf("res: %v", res)
}

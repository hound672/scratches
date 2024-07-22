package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

type Key [32]byte

func (k *Key) cryptBytes() []byte {
	return k[len(k)/2:]
}

func (k *Key) signBytes() []byte {
	return k[:len(k)/2]
}

// Generate initializes k with pseudorandom data from package crypto/rand.
func (k *Key) Generate() error {
	_, err := io.ReadFull(rand.Reader, k[:])
	return err
}

// Encode returns the URL-safe base64 encoding of k.
func (k *Key) Encode() string {
	return base64.URLEncoding.EncodeToString(k[:])
	// return hex.EncodeToString(k[:])
}

func main() {
	var key Key
	if err := key.Generate(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(key.Encode())
}
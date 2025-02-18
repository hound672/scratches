package main

import (
	"bytes"
	"io"
	"log"
)

func read1(r io.Reader) {
	d := make([]byte, 3)
	
	n, err := r.Read(d)
	if err != nil {
		panic(err)
	}

	log.Printf("Read %d bytes. Data: %v", n, string(d))
}

func read2(r io.Reader) {
	d := make([]byte, 1)
	
	n, err := r.Read(d)
	if err != nil {
		panic(err)
	}

	log.Printf("Read %d bytes. Data: %v", n, string(d))
}


func main() {
	log.Printf("Start")
	
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	read1(bytes.NewReader(data))
	read2(bytes.NewReader(data))
}

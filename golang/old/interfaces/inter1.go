package main

import "log"

type Reader interface {
	Read()
}

type OwnClass struct {
	Reader
}

func (o *OwnClass) Read() {
	log.Printf("read is called")
}

func read(i Reader) {
	log.Printf("called with reader: %v", i)
	i.Read()
}

func main() {
	log.Printf("Start")
	
	own := &OwnClass{}
	read(own)
}

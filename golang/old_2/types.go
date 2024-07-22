package main

import "log"

type Path string

func CP(f Path) {

}

func main() {
	log.Printf("hello world")

	p := Path("/tmp/file.txt")
	s := "sdsdf"
	CP(s)

	log.Printf("c: %v; type: %T", p, p)
}

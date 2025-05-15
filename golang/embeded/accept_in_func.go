package main

import (
	"log"
)

type SomeEmbed struct {
	Value int
}

type MainStruct struct {
	MainValue int
	SomeEmbed
}

func accept(e SomeEmbed) {
	log.Printf("Got: %v", e)
}

func main() {
	log.Printf("Start")

	e := SomeEmbed{}
	accept(e)

	m := MainStruct{SomeEmbed: e}
	accept(m.SomeEmbed)
}
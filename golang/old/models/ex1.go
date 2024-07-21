package main

import "log"

type File struct {
	Field string
}

type MainP struct {
	*File
}

type MainV struct {
	File
}

func main() {
	mP := &MainP{
		&File{
			"222",
		},
	}

	mV := &MainV{
		File{Field: "111"},
	}

	log.Printf("mp: %v; mv: %v", mP, mV)
}

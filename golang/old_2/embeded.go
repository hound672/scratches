package main

import "log"

type File struct {
	Name string
	Path string
}

type Artifact struct {
	*File
}

func fs(f *File) {
	log.Printf("fff: %v", f)
}

func main() {
	log.Printf("start")

	art := &Artifact{
		&File{Name: "name", Path: "path"},
	}
	log.Printf("art: %v", art.File)
	
	fs(art.File)
}

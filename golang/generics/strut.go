package main

import (
	"log"
)

type Repo struct {
	VCS string
}

type Cloud struct {
	URL string
}

type Assets interface{
	Repo | Cloud
}

type Asset[T Assets] struct {
	Name string
	Payload T
}

func Get() Asset[Repo] {
	return Asset[Repo]{
		Name: "repo",
		Payload: Repo{VCS: "git"},
	}
}

func main() {
	log.Printf("Start")
	
	asset := Asset[Repo]{
		Name: "test_asset",
	}
	
	log.Printf("Asset: %v", asset)
	
	log.Printf("From func %+v", Get())
	
}

package main

import (
	"log"
	"sort"
)

type Artifact struct {
	Version int
	Name string
}

type ArtifactSortable []Artifact

func (a ArtifactSortable) Len() int           { return len(a) }
func (a ArtifactSortable) Less(i, j int) bool { return a[i].Version < a[j].Version }
func (a ArtifactSortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	arts := []Artifact{
		{Version: 100, Name: "100"},
		{Version: 0, Name: "0"},
		{Version: 200, Name: "200"},
		{Version: 50, Name: "50"},
		{Version: 1, Name: "1"},
	}

	log.Printf("Before sort; %v", arts)

	sort.Sort(ArtifactSortable(arts))
	log.Printf("After sort: %v", arts)
	
	//sort.Sort(sort.Reverse(ArtifactSortable(arts)))
	log.Printf("After revr: %v", arts)
	toFind := 200
	
	r := sort.Search(len(arts), func(i int) bool {
		return arts[i].Version >= toFind
	})
	if r < len(arts) && arts[r].Version == toFind {
		log.Printf("FOUND: %v", r)
	} else {
		log.Printf("not found!!!")
	}
	
	//log.Printf("r: %v", r)
}

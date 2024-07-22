package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.Printf("Start")

	// source := "google/streetview/publish/v1/streetview_publish_gapic.yaml"
	source := "google/streetview_publish_gapic.yaml"
	// source := "streetview_publish_gapic.yaml"

	dir, file := filepath.Split(source)
	
	log.Printf("dir: %s; file: %s", dir, file)

	dirs := strings.Split(dir, string(os.PathSeparator))
	
	log.Printf("d: %s", dirs[0])
}

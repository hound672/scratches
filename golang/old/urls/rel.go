package main

import (
	"log"
	"net/url"
)

func main() {
	log.Printf("Start")

	u := "http://localhost:9000/link/one?val=123&param=456"
	urlLink, err := url.Parse(u)
	if err != nil {
		log.Fatalf("err parse url: %v", err)
	}
	log.Printf("urlLink: %v", urlLink.Path)
	log.Printf("str: %s", urlLink.String())
	log.Printf("urlLink: %v", urlLink.RequestURI())

	//relPath, err := filepath.Rel(url, "http://localhost:9000")
	//if err != nil {
	//	log.Fatalf("Err: %v", err)
	//}
	//log.Printf("rel path: %s", relPath)
}

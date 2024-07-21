package main

import ("log"
	"time"
)

func main() {
	log.Printf("Start")

	casdoorData := "2023-04-26T17:23:42+03:00"
	//layout := "2006-01-02T15:04:05.000Z"
	//layout := "2006-01-02T15:04:05"
	layout := time.RFC3339

	log.Printf("source date: %s", casdoorData)

	d, err := time.Parse(layout, casdoorData)
	log.Printf("d: %v; err: %v", d, err)
	
	toStr := d.Format(layout)
	log.Printf("after: %s", toStr)
}
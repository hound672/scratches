package main

import (
	"io"
	"log"
	"os"
	"time"
)

func main() {
	log.Printf("Start")

	fileName := "/tmp/easyp/easyp.lock"

	fp, err := os.OpenFile(fileName, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	res, err := fp.WriteString("Hello\n")
	log.Printf("res: %v; err: %v", res, err)
	// fp.Close()
	wr(fp)

	time.Sleep(time.Second * 10)
}

func wr(w io.Writer) {
	w.Write([]byte("Here is a string...."))
}

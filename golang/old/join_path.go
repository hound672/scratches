package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	ext := ".html"
	path := "/tmp/var"
	name := "email"

	res := filepath.Join(path, name + ext)
	log.Println(res)

	var f os.File
	var c net.Conn

	//f.Read()
	c.Read()

	_ = f
}

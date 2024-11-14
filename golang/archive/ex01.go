// Create archive in ram
package main

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"path"
)

func storeFile(r io.Reader) error {
	tmpDir := os.TempDir()
	tmpFileName := path.Join(tmpDir, "archive.zip")

	fp, err := os.OpenFile(tmpFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	if _, err := fp.Write(content); err != nil {
		panic(err)
	}
	log.Printf("file: %s", tmpFileName)

	return nil
}

func main() {
	log.Printf("Start")
	
	fileName := "test.txt"
	fileContent := "fome file content"

	buffer := &bytes.Buffer{}

	archive := zip.NewWriter(buffer)

	f, err := archive.Create(fileName)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(fileContent))
	if err != nil {
		panic(err)
	}
	if err := archive.Close(); err != nil {
		panic(err)
	}

	if err := storeFile(buffer); err != nil {
		panic(err)
	}
}

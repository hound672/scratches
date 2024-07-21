package main

import (
	"errors"
	"log"

	"google.golang.org/grpc/status"
)

func main() {
	log.Printf("Start")

	err := errors.New("some error")

	gErr := status.Convert(err)
	log.Printf("gErr: %v", gErr)
	log.Printf("gErr: %v", gErr.Code())

}

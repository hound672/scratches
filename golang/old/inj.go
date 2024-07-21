package main

import (
	"log"
)

type Worker interface {
	Do() error
}

func Do() error {
	log.Printf("Do is called from func")
	return nil
}


//

type worker struct {}

func (w *worker) Do() error {
	log.Printf("Do is called from struct")
	return nil
}

//


func Handler(worker Worker) {
	_ = worker.Do()
}

func main() {
	log.Printf("start")

	Handler(worker{})
}
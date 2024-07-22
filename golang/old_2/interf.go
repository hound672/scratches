package main

import "log"

type AbsRepo interface {
	Read() error
}

type ImpRepo struct {

}

func (r *ImpRepo) Read() error {
	return nil
}

func main() {
	var repo AbsRepo
	repo = &ImpRepo{}

	_ = repo

	log.Print("start")
}

package main

import (
	"errors"
	"log"
)

var (
	ErrOwn = errors.New("some error")
)

func getErr() error {
	return ErrOwn
	//return errors.New("other error")
	//return nil
}

func doIf() {
	err := getErr()

	if err != nil {
		if errors.Is(err, ErrOwn) {
			//log.Printf("got errOwn")
			return
		}
		//log.Printf("got other error")
	} else {
		//log.Printf("not error")
	}
}

func doCase() {
	err := getErr()

	if err != nil {
		switch {
		case errors.Is(err, ErrOwn):
			//log.Printf("got errOwn")
		default:
			//log.Printf("got other error")
		}

	}

	//switch {
	//case errors.Is(err, ErrOwn):
	//	//log.Printf("got errOwn")
	//case err == nil:
	//	//log.Printf("not error")
	//default:
	//	//log.Printf("got other error")
	//}
}


func main() {
	log.Printf("Run err bench")

	doIf()
	doCase()
}


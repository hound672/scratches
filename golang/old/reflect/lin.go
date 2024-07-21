package main

import (
	"log"
	"reflect"
)

type Rule interface{
	Validate() error
}

type LinterRule struct {}

func (l LinterRule) Validate() error {
	log.Printf("Called validate in LinterRule")
	return nil
}

type CommentRule struct{

}

func (l *CommentRule) Validate() error {
	return nil
}

func main() {
	log.Printf("start")

	//var r Rule = &LinterRule{}
	var r Rule = &CommentRule{}

	_ = r

	name := reflect.TypeOf(r).Elem().Name()
	log.Printf("name: %s", name)
}

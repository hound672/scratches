package main

import (
	"log"
)

type AuthProfiles interface {
	Auth1 | Auth2
}

type Auth1 struct {
	Username string
}

type Auth2 struct {
	Token string
}

type Profile struct {
	Profile interface{}
}

func Get[T AuthProfiles](profile *Profile) T {
	obj, ok := profile.Profile.(T)
	if !ok {
		panic("!!!")
	}
	
	return obj
}

func main() {
	auth1 := Auth2{
		// Username: "test",
		Token: "token",
	}
	
	profile := &Profile{
		Profile: auth1,
	}
	
	res := Get[Auth2](profile)
	// log.Printf("res: %v; field: %v, type: %T", res, res.Username, res)
	log.Printf("res: %v; field: %v, type: %T", res, res.Token, res)
}

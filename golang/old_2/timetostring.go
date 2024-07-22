package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type OwnTime time.Time

func (o OwnTime) String() string {
	return fmt.Sprintf("!!!!!!: %s", time.Time(o).Format(time.RFC3339))
}

func (o OwnTime) MarshalJSON() ([]byte, error) {
	log.Printf("MARRR")
	return []byte(`"HELLO"`), nil
}

func main() {
	log.Printf("start")

	t := OwnTime(time.Now())
	
	type S struct {
		Time OwnTime `json:"time"`
		//Time time.Time `json:"time"`
	}
	log.Printf("t: %s", t)
	
	s := &S{Time: t}
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("err: %v", err)
	} 
	
	log.Printf("res: %v", string(res))
}

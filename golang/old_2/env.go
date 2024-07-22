package main

import (
	"log"
	"strings"
)

func main() {
	log.Printf("Start")

	source := []byte(`
KEY_1=VAL_1
KEY_2=VAL_2
KEY_3=VAL_3
KEY_4=
KEY_5
`)

	res := map[string]string{}

	r := string(source)

	for _, v := range strings.Split(r, "\n") {
		if v == "" {
			continue
		}

		parsed := strings.Split(v, "=")
		log.Printf("v: %v", parsed)
		kr := parsed[0]
		var vr string
		if len(parsed) == 2 {
			vr = parsed[1]
		} else {
			vr = ""
		}

		res[kr] = vr
	}

	log.Printf("res: %v", res)
}

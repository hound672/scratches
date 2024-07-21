package main

import "log"

func main() {
	log.Printf("Start")
	
	m := make(map[string]string)
	m["key"] = "dfdsf"
	
	p, ok := m["key"]
	log.Printf("p: %v; ok: %v", p, ok)
	
	var field string
	
	field = p
	
	log.Printf("field: %s", field)
}

package main

import (
	"log"
	"strings"
)

type instructionInfo struct {
	pkgName     string
	instruction string
}

func main() {
	log.Printf("Start get_package")

	// source := "google.api.http"
	// source := "grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation"
	// source := "StartScanRequest"
	// source := "buf.validate.field.string.uuid"
	source := "(buf.validate.field).string.uuid"
	// source := "(google.api.http)"

	res := parseInstruction(source)

	log.Printf("res: %v", res)

	//

	log.Printf("start another")

	m := make(map[string][]string)
	m["test"] = []string{"111", "222"}

	for i, v := range m["test"] {
		log.Printf("i: %v; v: %v", i, v)
	}
}

// parseInstruction parse input string and return its package name
// return empty string if passed input does not imported from another package
func parseInstruction(input string) instructionInfo {
	// check if there is brackets, and extract
	// (google.api.http) -> google.api.http
	// (buf.validate.field).string.uuid -> buf.validate.field
	// or pkg.FieldType -> pkg.FieldType
	iStart := strings.Index(input, "(")
	iEnd := strings.Index(input, ")")
	if iStart != -1 && iEnd != -1 {
		input = input[iStart+1 : iEnd]
	}

	idx := strings.LastIndex(input, ".")
	if idx <= 0 {
		return instructionInfo{instruction: input}
	}

	return instructionInfo{
		pkgName:     input[:idx],
		instruction: input[idx+1:],
	}
}

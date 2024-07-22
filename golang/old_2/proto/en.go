package main

import (
	"log"

	bbAPI "gitlab.ptsecurity.com/appsec-platform/contracts-proto/golang/cross_cases/smart_start/ptbb_integration_grpc_api/v1"
)

func main() {
	log.Printf("Start")

	res := bbAPI.StartScanErrorCode_NO_ERROR.String()
	log.Printf("res: %v", res)
}

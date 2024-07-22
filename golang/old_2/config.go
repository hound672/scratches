package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("Start")

	address := "ya.ru"
	port := 1234
	res := fmt.Sprintf("http://%s:%d", address, port)

	fmt.Printf("res: %v\n", res)

	res = net.JoinHostPort(address, string(port))
	fmt.Printf("res2: %v\n", res)
}

package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	//cmd := exec.Command("docker", "version", "--format", "{{.Server.Version}}")
	cmd := exec.Command("docker-compose", "version", "--format", "{{.Server.Version}}")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	apiVersion := strings.TrimSpace(string(cmdOutput.Bytes()))
	
	fmt.Printf("version: %v\n", apiVersion)


	cmd = exec.Command("python", "/tmp/checkl/s.py")
	cmdOutput = &bytes.Buffer{}
	cmdFail := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdFail

	err = cmd.Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		//panic(err)
	}

	fmt.Printf("out: %v\n", cmdOutput)
	fmt.Printf("fail: %v\n", cmdFail)
}
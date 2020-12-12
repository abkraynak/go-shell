package main

import (
	"fmt"
	"strings"
	"os"
	"os/exec"
)

func main() {
	for {
		var line string
		fmt.Print(">> ")
		fmt.Scanln(&line)
		exec_cmd(line)
	}
}

func exec_cmd(in string) error {
	// remove newline and separate string
	in = strings.TrimSuffix(in, "\n")
	args := strings.Split(in, " ")
	
	switch args[0] {
		case "exit":
			os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)	
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

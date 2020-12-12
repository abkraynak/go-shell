package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		scanner.Scan()
		line := scanner.Text()
		execCmd(line)
	}
}

func execCmd(in string) (error) {
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

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		return err		
	}
	return nil
}


package main

import (
	"fmt"
)

func main() {
	for {
		var line string
		fmt.Scanln(&line)
		fmt.Println("You entered " + line)
	}
}


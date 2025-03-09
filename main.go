package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		return
	}

	fmt.Println("Running command:", os.Args[1])
}

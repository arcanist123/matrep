package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	fileNamePtr := flag.String("file", "", "file to parse as a matrix")

	// Parse the flags from the command line
	flag.Parse()

	// Access the values of the flags
	fileName := *fileNamePtr

	fmt.Printf("Name: %s\n", fileName)
	if fileName != "" {

	}

	// Accessing non-flag arguments (after the flags)
	fmt.Println("Non-flag arguments:", flag.Args())
	fmt.Println("Number of non-flag arguments:", flag.NArg())
}

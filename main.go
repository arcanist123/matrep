package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	fileNamePtr := flag.String("file", "data.csv", "file to parse as a matrix")
	configFilePtr := flag.String("config_name", "config1", "name of configuration")

	// Parse the flags from the command line
	flag.Parse()

	// Access the values of the flags
	fileName := *fileNamePtr
	configFile := *configFilePtr
	fmt.Printf("Name: %s\n", fileName)
	fmt.Printf("Config: %s\n", configFile)
	if fileName == "" || configFile == "" {
		fmt.Printf("cannot proceed with processing")

	} else {
		// var config =
	}

}

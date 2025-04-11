package main

import (
	"flag"
	"fmt"
)

func main() {
	test_yaml()
	// Define flags
	fileNamePtr := flag.String("file", "data.csv", "file to parse as a matrix")
	configNamePtr := flag.String("config_name", "config1", "name of configuration")

	// Parse the flags from the command line
	flag.Parse()

	// Access the values of the flags
	fileName := *fileNamePtr
	configName := *configNamePtr
	fmt.Printf("Name: %s\n", fileName)
	fmt.Printf("Config: %s\n", configName)
	if fileName == "" || configName == "" {
		fmt.Printf("cannot proceed with processing")

	} else {
		aaa, err := NewConfigFactory(configName).GetConfig()
		if err != nil {
			fmt.Println("config initialisation failed")
		}
		fmt.Println(aaa)
	}

}

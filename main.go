package main

import (
	"flag"
	"fmt"

	"github.com/arcanist123/matrep/config"
	"github.com/arcanist123/matrep/engine"
)

func main() {
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
		config, err := config.NewConfigFactory(configName).GetConfig()
		if err != nil {
			fmt.Printf("error loading config: %s\n", err)
		}
		matrix, err := NewFileReader(fileName).readCSVToMatrix()
		if err != nil {
			fmt.Printf("error reading file: %s\n", err)
		}

		resolvedMatrix := engine.NewMatrixHandler(matrix, config)
		fmt.Println(resolvedMatrix)
	}

}

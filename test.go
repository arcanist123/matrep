package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3" // Using the recommended v3 version
)

// Define a sample structure
type Person struct {
	Name    string   `yaml:"name"`
	Age     int      `yaml:"age"`
	Hobbies []string `yaml:"hobbies,omitempty"` // omitempty will skip if the slice is empty
}

func test_yaml() {
	// Create a slice of Person structures
	people := []Person{
		{
			Name:    "Alice",
			Age:     30,
			Hobbies: []string{"Reading", "Hiking"},
		},
		{
			Name: "Bob",
			Age:  25,
		},
		{
			Name:    "Charlie",
			Age:     35,
			Hobbies: []string{"Coding", "Gaming", "Music"},
		},
	}

	// Convert the slice of structures to YAML
	yamlData, err := yaml.Marshal(&people)
	if err != nil {
		log.Fatalf("Error marshaling YAML: %v", err)
	}

	// Print the YAML output
	fmt.Println(string(yamlData))
}

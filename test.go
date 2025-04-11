package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type Config1 struct {
	Name       string `yaml:"name"`
	ConfigType string `yaml:"configType"`
	Systems    string `yaml:"systems"`
}

func test() {
	yamlData := `
- name: "config1"
  configType: "file"
  systems: systems.csv
- name: "config2"
  configType: "db"
  systems: ZSYSTEMS
`

	var configs []Config1

	err := yaml.Unmarshal([]byte(yamlData), &configs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("%+v\n", configs)
}

package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigFactory struct {
	configName string
}

func NewConfigFactory(configName string) ConfigFactory {
	return ConfigFactory{configName: configName}
}
func (me ConfigFactory) GetConfig() (result Config, err error) {

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	fmt.Println(string(yamlFile))
	var configs []Config
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return
	}

	for _, config := range configs {
		if config.name == me.configName {
			return config, nil
		}
	}

	return Config{}, nil
}

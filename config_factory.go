package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	fileName string
}

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func NewConfig(fileName string) Config {
	return Config{fileName: fileName}
}

func (this Config) GetConfig() (result string, err error) {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return
	}

	fmt.Printf("Server Host: %s\n", config.Server.Host)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("Database Username: %s\n", config.Database.Username)
	fmt.Printf("Database Password: %s\n", config.Database.Password)

}

package main

type Config struct {
	name       string `yaml:"name"`
	configType string `yaml:"configType"`
	systems    string `yaml:"systems"`
}

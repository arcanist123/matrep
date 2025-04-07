package main

type Config struct {
	fileName string
}

func NewConfig(fileName string) Config {
	return Config{fileName: fileName}
}

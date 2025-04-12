package config

import (
	"errors"
	"fmt"
	"log"
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

	var data any
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
	switch v := data.(type) {
	case []any:
		for _, item := range v {
			if configMap, ok := item.(map[string]any); ok { // Type assertion here
				result, err = me._parseConfig(configMap)
				if err == nil {
					return result, nil
				}
			} else {
				return Config{}, errors.New("unexpected structure of yaml file")
			}
		}
	default:
		return Config{}, errors.New("unexpected structure of yaml file")
	}

	return Config{}, errors.New("requested config was not found")
}

func (me ConfigFactory) _parseConfig(configMap map[string]any) (result Config, err error) {
	result = Config{}
	for key, value := range configMap {
		switch key {
		case "name":
			if name, ok := value.(string); ok {
				result.name = name
			} else {
				return Config{}, fmt.Errorf("expected string for key 'name', got %T", value)
			}
		case "configType":
			if file, ok := value.(string); ok {
				result.configType = file
			} else {
				return Config{}, fmt.Errorf("expected string for key 'file', got %T", value)
			}
		case "systems":
			if systems, ok := value.(string); ok {
				result.systems = systems
			} else {
				return Config{}, fmt.Errorf("expected string for key 'file', got %T", value)
			}
		}
	}

	if result.name != me.configName {
		return Config{}, fmt.Errorf("this is not the requested config")
	} else {
		return result, nil
	}
}

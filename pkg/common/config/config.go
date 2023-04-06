package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func LoadConfig(config []byte, configuration interface{}) (interface{}, error) {
	// unmarshal the config file
	err := yaml.Unmarshal(config, configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %v", err)
	}

	return configuration, nil
}

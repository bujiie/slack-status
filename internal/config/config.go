package config

import (
	"context"
	"gopkg.in/yaml.v3"
	"os"
)

const UnknownChar = "?"

type YamlConfig struct {
	Mapping map[string]string `yaml:"mapping"`
}

func (y YamlConfig) Get(char string) string {
	if val, ok := y.Mapping[char]; ok {
		return val
	}
	return UnknownChar
}

func ParseConfig(_ context.Context, filename string) (*YamlConfig, error) {
	fp, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var yamlConfig YamlConfig
	if err := yaml.Unmarshal(fp, &yamlConfig); err != nil {
		return nil, err
	}
	return &yamlConfig, nil
}

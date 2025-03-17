package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var DefaultFileName = ".status_config"

type Config struct {
	Prefix        string            `yaml:"prefix"`
	Start         string            `yaml:"start" default:"Sunday"`
	IgnoreWeekend bool              `yaml:"ignoreWeekend"`
	IncDayOfWeek  bool              `yaml:"incDayOfWeek"`
	Symbols       map[string]string `yaml:"symbols"`
}

func ReadYAML(filepath string, out interface{}) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, out)
}

type ContextKey string

const (
	MomentKey             ContextKey = "moment"
	StartOfWeekKey        ContextKey = "startofweek"
	IncDayOfWeekPrefixKey ContextKey = "incdayofweekprefix"
	IgnoreWeekendKey      ContextKey = "ignoreweekend"
)

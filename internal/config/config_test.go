package config

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharMappingConfig(t *testing.T) {
	t.Run("should read mapping from config file and hydrate into an object", func(t *testing.T) {
		config, _ := ParseConfig(context.Background(), "./config_test.yaml")
		assert.Contains(t, config.Mapping, "a")
		assert.Contains(t, config.Mapping, "b")
		assert.Contains(t, config.Mapping, "c")
		assert.Equal(t, "aye", config.Mapping["a"])
		assert.Equal(t, "bee", config.Mapping["b"])
		assert.Equal(t, "see", config.Mapping["c"])
	})
}

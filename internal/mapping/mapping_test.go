package mapping

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const TestUnknownChar = "?"

type TestMappingType map[string]string

func (tm TestMappingType) Get(char string) string {
	if val, ok := tm[char]; ok {
		return val
	}
	return TestUnknownChar
}

func TestMapping(t *testing.T) {
	var testMapping CharMapping = TestMappingType{
		"s": "r",
	}

	t.Run("should return a pre-mapped value associated with the specified character", func(t *testing.T) {
		subject := "s"
		assert.Equal(t, "r", testMapping.Get(subject))
	})

	t.Run("should return the unknown placeholder when the specified character does not have a mapping", func(t *testing.T) {
		subject := "u"
		assert.Equal(t, "?", testMapping.Get(subject))
	})
}

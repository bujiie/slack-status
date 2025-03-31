package color

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	t.Run("should return original string between color ASCII tags", func(t *testing.T) {
		subject := "I am some text"
		result := Colorize(subject, Yellow)
		assert.True(t, strings.HasPrefix(result, Yellow))
	})
}

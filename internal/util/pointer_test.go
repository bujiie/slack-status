package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToPointer(t *testing.T) {
	t.Run("should return a pointer to the original integer", func(t *testing.T) {
		subject := 123
		intPointer := IntToPointer(subject)
		assert.NotNil(t, intPointer)
		assert.Equal(t, subject, *intPointer)
	})
}

func TestStrToPointer(t *testing.T) {
	t.Run("should return a pointer to the original string", func(t *testing.T) {
		subject := "original string"
		strPointer := StrToPointer(subject)
		assert.NotNil(t, strPointer)
		assert.Equal(t, subject, *strPointer)
	})
}

package mapping

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapping(t *testing.T) {
	ctx := context.Background()
	testMapping := CharMapping{
		"s": "r",
	}

	t.Run("should return a pre-mapped value associated with the specified character", func(t *testing.T) {
		subject := "s"
		assert.Equal(t, "r", testMapping.GetMapping(ctx, subject))
	})

	t.Run("should return the unknown placeholder when the specified character does not have a mapping", func(t *testing.T) {
		subject := "u"
		assert.Equal(t, "?", testMapping.GetMapping(ctx, subject))
	})
}

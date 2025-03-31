package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	tz, _ := time.LoadLocation("America/Los_Angeles")
	targetDate := time.Date(2000, 1, 1, 0, 0, 0, 0, tz)

	t.Run("should return an error message if offset is not a number", func(t *testing.T) {
		_, err := status(context.Background(), targetDate, []string{"notANumber"}...)
		assert.ErrorContains(t, err, "Error: cannot make status because offset specified does not look like a number")
	})
}

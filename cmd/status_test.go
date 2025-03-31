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
		_, err := status(context.Background(), targetDate, []string{"notANumber", ""}...)
		assert.ErrorContains(t, err, "Error: cannot make status because offset specified does not look like a number")
	})

	t.Run("should return the week number for the specified date", func(t *testing.T) {
		result, _ := status(context.Background(), targetDate, []string{""}...)
		assert.Equal(t, "week 1: ", *result)
	})

	t.Run("should return the week number for the specified date and positive offset", func(t *testing.T) {
		result, _ := status(context.Background(), targetDate, []string{"+2", ""}...)
		assert.Equal(t, "week 3: ", *result)
	})

	t.Run("should return the week number for the specified date and negative offset", func(t *testing.T) {
		result, _ := status(context.Background(), targetDate, []string{"-2", ""}...)
		assert.Equal(t, "week 51: ", *result)
	})

	t.Run("should return the week number for the specified date and offset that moves target date into next year", func(t *testing.T) {
		result, _ := status(context.Background(), targetDate, []string{"+53", ""}...)
		assert.Equal(t, "week 1: ", *result)
	})

	t.Run("should return the week number for the specified date and offset that moves target date into previous year", func(t *testing.T) {
		result, _ := status(context.Background(), targetDate, []string{"-52", ""}...)
		assert.Equal(t, "week 1: ", *result)
	})

	t.Run("should return the week number for the specified pattern", func(t *testing.T) {
		result, _ := status(context.Background(), targetDate, []string{"oh"}...)
		assert.Equal(t, "week 1: :building::house:", *result)
	})

	t.Run("should return the status for the specified date offset and pattern", func(t *testing.T) {
		result, _ := status(context.Background(), targetDate, []string{"+1", "oh"}...)
		assert.Equal(t, "week 2: :building::house:", *result)
	})
}

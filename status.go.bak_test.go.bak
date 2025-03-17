package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetWeekNumber(t *testing.T) {
	pst, _ := time.LoadLocation("America/Los_Angeles")

	t.Run("should identify the first day of the year to be in week 1", func(t *testing.T) {
		testMoment := time.Date(2025, time.January, 1, 0, 0, 0, 0, pst) // Wednesday
		result := getWeekNumber(testMoment, time.Sunday)
		assert.Equal(t, 1, result)
	})

	t.Run("should identify the the first Wednesday of the year to be in week 1 with starting day on Friday", func(t *testing.T) {
		testMoment := time.Date(2025, time.January, 1, 0, 0, 0, 0, pst) // Wednesday
		result := getWeekNumber(testMoment, time.Friday)
		assert.Equal(t, 1, result)
	})

	t.Run("should", func(t *testing.T) {
		testMoment := time.Date(2025, time.January, 6, 0, 0, 0, 0, pst) // Wednesday
		result := getWeekNumber(testMoment, time.Tuesday)
		assert.Equal(t, 1, result)
	})
}

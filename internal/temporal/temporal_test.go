package temporal

import (
	"context"
	"github.com/bujiie/slack-status/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetWeekNumber(t *testing.T) {
	tz, _ := time.LoadLocation("America/Los_Angeles")
	targetDate := time.Date(2000, 1, 1, 0, 0, 0, 0, tz)

	t.Run("should return the week number of the year for the given moment", func(t *testing.T) {
		weekNumber := GetWeekNumber(context.Background(), targetDate, nil)
		assert.Equal(t, 1, weekNumber)
	})

	t.Run("should return the week number of the year when the offset moves the moment into the previous year", func(t *testing.T) {
		weekNumber := GetWeekNumber(context.Background(), targetDate, util.IntToPointer(-2))
		assert.Equal(t, 51, weekNumber)
	})

	t.Run("should return the week number of the year when the offset moves the moment into the next year", func(t *testing.T) {
		weekNumber := GetWeekNumber(context.Background(), targetDate, util.IntToPointer(53))
		assert.Equal(t, 1, weekNumber)
	})

	t.Run("should return the week number of the year for two weeks into the future", func(t *testing.T) {
		weekNumber := GetWeekNumber(context.Background(), targetDate, util.IntToPointer(2))
		assert.Equal(t, 3, weekNumber)
	})

	t.Run("should return the week number of the year for 2 weeks into the past", func(t *testing.T) {
		weekNumber := GetWeekNumber(context.Background(), targetDate, util.IntToPointer(-2))
		assert.Equal(t, 51, weekNumber)
	})
}

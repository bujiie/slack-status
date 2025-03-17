package temporal

import (
	"context"
	"github.com/bujiie/slack-status/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetSTartOfWeek(t *testing.T) {
	tz, _ := time.LoadLocation("America/Los_Angeles")
	targetDate := time.Date(2000, 1, 1, 0, 0, 0, 0, tz)

	t.Run("should return the date for the Sunday of the week that includes January 1, 2000", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, config.MomentKey, targetDate)
		ctx = context.WithValue(ctx, config.StartOfWeekKey, time.Sunday)

		startOfWeekTime := GetStartOfWeek(ctx)
		// Jan 1, 2000 falls on a Saturday, so the week that starts on a Sunday
		// that includes this Saturday is Dec. 26, 1999.
		assert.Equal(t, 1999, startOfWeekTime.Year())
		assert.Equal(t, time.Month(12), startOfWeekTime.Month())
		assert.Equal(t, 26, startOfWeekTime.Day())
	})

	t.Run("should return the same date when the start day of week is the same date", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, config.MomentKey, targetDate)
		ctx = context.WithValue(ctx, config.StartOfWeekKey, time.Saturday)

		startOfWeekTime := GetStartOfWeek(ctx)
		assert.Equal(t, 2000, startOfWeekTime.Year())
		assert.Equal(t, time.Month(1), startOfWeekTime.Month())
		assert.Equal(t, 1, startOfWeekTime.Day())
	})
}

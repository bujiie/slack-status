package temporal

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetWeekNumber(t *testing.T) {
	tz, _ := time.LoadLocation("America/Los_Angeles")
	targetDate := time.Date(2000, 1, 1, 0, 0, 0, 0, tz)

	t.Run("should return the week number of the year for the given moment", func(t *testing.T) {
		weekNumber := GetWeekNumber(context.Background(), targetDate)
		assert.Equal(t, 1, weekNumber)
	})
}

package temporal

import (
	"context"
	"github.com/bujiie/slack-status/internal/config"
	"time"
)

// GetWeekNumber calculates the week number out of the year that contains the
// time 't' where the start of the week is defined by 'startOfWeek'.
func GetWeekNumber(ctx context.Context) int {
	moment := ctx.Value(config.MomentKey).(time.Time)
	startOfWeek := ctx.Value(config.StartOfWeekKey).(time.Weekday)
	// get the beginning of the year that includes time 't' in the timezone of
	// time 't'.
	startOfYear := time.Date(moment.Year(), time.January, 1, 0, 0, 0, 0, moment.Location())

	// calculate how many days BEFORE Jan. 1 to start counting as the beginning
	// of week 1.
	dateOffset := int(startOfWeek) - int(startOfYear.Weekday())
	// positive offset indicates the 'startOfWeek' occurs AFTER Jan. 1. We need
	// to subtract a week's worth of days (7) to maintain the same day of week
	// while remaining BEFORE Jan. 1.
	if dateOffset > 0 {
		dateOffset -= 7
	}

	// should start on the closest day of the week that matches 'startOfWeek'
	// that occurs BEFORE Jan. 1 ensuring Jan. 1 is in week one.
	adjustedStartOfYear := startOfYear.AddDate(0, 0, dateOffset)

	//  elapsed time since time 't' since the beginning of week one
	elapsedHours := moment.Sub(adjustedStartOfYear).Hours()
	elapsedDays := int(elapsedHours / 24)
	elapsedWeeks := elapsedDays / 7

	return elapsedWeeks + 1
}

func GetStartOfWeek(ctx context.Context) time.Time {
	moment := ctx.Value(config.MomentKey).(time.Time)
	startOfWeek := ctx.Value(config.StartOfWeekKey).(time.Weekday)

	daysSinceStart := (7 + int(moment.Weekday()) - int(startOfWeek)) % 7
	startOfWeekTime := moment.AddDate(0, 0, -daysSinceStart)
	return time.Date(startOfWeekTime.Year(), startOfWeekTime.Month(), startOfWeekTime.Day(), 0, 0, 0, 0, moment.Location())
}

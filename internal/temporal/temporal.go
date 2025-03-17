package temporal

import (
	"context"
	"github.com/bujiie/slack-status/internal/config"
	"strings"
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

// GetStartOfWeek returns the date for the start day of the week that contains
// the context moment. Note that this function takes into account the
// configured start day of the week.
func GetStartOfWeek(ctx context.Context) time.Time {
	moment := ctx.Value(config.MomentKey).(time.Time)
	startOfWeek := ctx.Value(config.StartOfWeekKey).(time.Weekday)

	//ctxDayOfWeek := int(moment.Weekday())
	//cfgStartOfWeek := int(startOfWeek)

	daysSinceStart := (7 + int(moment.Weekday()) - int(startOfWeek)) % 7
	startOfWeekTime := moment.AddDate(0, 0, -daysSinceStart)
	return time.Date(startOfWeekTime.Year(), startOfWeekTime.Month(), startOfWeekTime.Day(), 0, 0, 0, 0, moment.Location())
}

// GetDayOfWeek converts a string representation of the day of the week with the
// corresponding enum value from the time package. Note that this function
// supports three forms of the day names, full, three, and two character forms.
func GetDayOfWeek(value string) time.Weekday {
	switch strings.ToLower(value) {
	case "sunday", "sun", "su":
		return time.Sunday
	case "monday", "mon", "mo":
		return time.Monday
	case "tuesday", "tue", "tu":
		return time.Tuesday
	case "wednesday", "wed", "we":
		return time.Wednesday
	case "thursday", "thu", "th":
		return time.Thursday
	case "friday", "fri", "fr":
		return time.Friday
	case "saturday", "sat", "sa":
		return time.Saturday
	default:
		return time.Sunday
	}
}

package temporal

import (
	"context"
	"time"
)

// GetWeekNumber calculates the week number out of the year that contains the
// time 't' where the start of the week is defined by 'startOfWeek'.
func GetWeekNumber(_ context.Context, moment time.Time, offset *int) int {
	if offset != nil {
		moment = moment.AddDate(0, 0, 7*(*offset))
	}

	// get the beginning of the year that includes time 't' in the timezone of
	// time 't'.
	startOfYear := time.Date(moment.Year(), time.January, 1, 0, 0, 0, 0, moment.Location())

	// calculate how many days BEFORE Jan. 1 to start counting as the beginning
	// of week 1.
	dateOffset := int(time.Sunday) - int(startOfYear.Weekday())
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

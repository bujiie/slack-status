package main

import (
	"fmt"
	"time"
)

// getWeekNumber calculates the week number out of the year that contains the
// time 't' where the start of the week is defined by 'startOfWeek'.
func getWeekNumber(t time.Time, startOfWeek time.Weekday) int {
	// get the beginning of the year that includes time 't' in the timezone of
	// time 't'.
	startOfYear := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())

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
	elapsedHours := t.Sub(adjustedStartOfYear).Hours()
	elapsedDays := int(elapsedHours / 24)
	elapsedWeeks := elapsedDays / 7

	return elapsedWeeks + 1
}

type SymbolLookup map[rune]string

func main() {
	now := time.Now()        // allow prev, next, +n, -n
	weekStart := time.Sunday // configurable

	weekNumber := getWeekNumber(now, weekStart)
	result := ""
	for _, char := range "ppohx" {
		result += string(charToEmoji(char))
	}
	fmt.Printf("wk%d: %s", weekNumber, result)
}

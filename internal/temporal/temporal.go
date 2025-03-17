package temporal

import (
	"strings"
	"time"
)

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

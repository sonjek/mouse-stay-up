package utils

import (
	"log"
	"strings"
	"time"
)

// IsInWorkingHours reports whether now falls within the given time window
// (format "HH:MM-HH:MM"). Equal boundaries mean "all day".
func IsInWorkingHours(now time.Time, timeWindow string) bool {
	parts := strings.SplitN(timeWindow, "-", 2)
	if len(parts) != 2 {
		log.Printf("utils: invalid time window format: %q", timeWindow)
		return false
	}

	if parts[0] == parts[1] {
		return true
	}

	start, ok := parseTimeToMinutes(parts[0])
	if !ok {
		log.Printf("utils: invalid start time: %q", parts[0])
		return false
	}

	end, ok := parseTimeToMinutes(parts[1])
	if !ok {
		log.Printf("utils: invalid end time: %q", parts[1])
		return false
	}

	return minutesSinceMidnight(now) >= start && minutesSinceMidnight(now) < end
}

// minutesSinceMidnight returns the number of minutes elapsed since midnight
// for the given time.
func minutesSinceMidnight(t time.Time) int {
	return t.Hour()*60 + t.Minute()
}

// parseTimeToMinutes parses "HH:MM" and returns minutes since midnight.
func parseTimeToMinutes(s string) (int, bool) {
	t, err := time.Parse("15:04", s)
	if err != nil {
		return 0, false
	}
	return minutesSinceMidnight(t), true
}

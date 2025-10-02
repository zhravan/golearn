package time

import "time"

// FormatDate formats a time.Time into "YYYY-MM-DD".
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDateTime formats a time.Time into "YYYY-MM-DD HH:MM:SS".
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// DaysBetween returns the number of days between two dates.
// Positive if end > start, negative if end < start.
func DaysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

// AddDays adds a specified number of days to a date and returns the new date.
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// ParseDate parses a string in "YYYY-MM-DD" format into a time.Time.
// Returns zero time if parsing fails.
func ParseDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}
	}
	return t
}

// ParseDateTime parses a string in "YYYY-MM-DD HH:MM:SS" format into a time.Time.
// Returns zero time if parsing fails.
func ParseDateTime(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return time.Time{}
	}
	return t
}

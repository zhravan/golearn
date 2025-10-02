package timeformatting

import (
	"time"
)

// FormatCurrentTime returns the current time formatted as "2006-01-02 15:04:05"
func FormatCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// ParseTimeString parses a time string with the layout "2006-01-02 15:04:05"
// and returns the parsed time and any error
func ParseTimeString(timeStr string) (any, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return nil, err
	}
	return parsedTime, nil
}

// FormatTime takes a time.Time and formats it using the given layout
func FormatTime(t any, layout string) string {
	if timeVal, ok := t.(time.Time); ok {
		return timeVal.Format(layout)
	}
	return ""
}

// ParseAndReformat parses a time string with inputLayout and reformats it with outputLayout
func ParseAndReformat(timeStr, inputLayout, outputLayout string) (string, error) {
	parsedTime, err := time.Parse(inputLayout, timeStr)
	if err != nil {
		return "", err
	}
	return parsedTime.Format(outputLayout), nil
}

// FormatWithTimezone formats the current time in the specified timezone
func FormatWithTimezone(timezone string) (string, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	return time.Now().In(loc).Format("2006-01-02 15:04:05 MST"), nil
}

// GetTimeComponents returns the year, month, day, hour, minute, second of a given time
func GetTimeComponents(t any) (int, int, int, int, int, int) {
	if timeVal, ok := t.(time.Time); ok {
		year, month, day := timeVal.Date()
		hour, minute, second := timeVal.Clock()
		return year, int(month), day, hour, minute, second
	}
	return 0, 0, 0, 0, 0, 0
}
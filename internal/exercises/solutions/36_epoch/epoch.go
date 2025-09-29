package epoch

import (
	"time"
)

// EpochToTime converts a unix epoch (seconds) into a formatted time string "2006-01-02 15:04:05".
func EpochToTime(epoch int64) string {
	// Convert the epoch to a time.Time object and set it to UTC before formatting.
	t := time.Unix(epoch, 0).UTC()
	return t.Format("2006-01-02 15:04:05")
}

// TimeToEpoch converts a formatted time string "2006-01-02 15:04:05" into a unix epoch (seconds).
// It assumes the input string represents a time in UTC.
func TimeToEpoch(input string) int64 {
	// Use time.ParseInLocation to explicitly parse the string as a UTC time.
	t, err := time.ParseInLocation("2006-01-02 15:04:05", input, time.UTC)
	if err != nil {
		// In a real application, you might log the error here.
		return 0
	}
	// t.Unix() returns the Unix epoch seconds.
	return t.Unix()
}
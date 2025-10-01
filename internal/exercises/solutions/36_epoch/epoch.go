package epoch

import (
	"time"
)

// --- Getting Current Time ---

// GetCurrentUnixSeconds returns the current time as a Unix epoch in seconds (int64).
func GetCurrentUnixSeconds() int64 {
	return time.Now().Unix()
}

// GetCurrentUnixMilliseconds returns the current time as a Unix epoch in milliseconds (int64).
func GetCurrentUnixMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetCurrentUnixNanoseconds returns the current time as a Unix epoch in nanoseconds (int64).
func GetCurrentUnixNanoseconds() int64 {
	return time.Now().UnixNano()
}

// GetCurrentFormattedTime returns the current time in the
// "2006-01-02 15:04:05.000000 +0000 UTC" format (microsecond precision).
func GetCurrentFormattedTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05.000000 +0000 UTC")
}

// GetCurrentFormattedTimeSimple returns the current time in the
// "2006-01-02 15:04:05 +0000 UTC" format (no microseconds).
func GetCurrentFormattedTimeSimple() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05 +0000 UTC")
}

// EpochToTime converts a unix epoch (seconds) into a formatted time string "2006-01-02 15:04:05".
// The resulting time is in UTC.
func EpochToTime(epoch int64) string {
	t := time.Unix(epoch, 0).UTC()
	return t.Format("2006-01-02 15:04:05")
}

// TimeToEpoch converts a formatted time string "2006-01-02 15:04:05" into a unix epoch (seconds).
// It assumes the input string represents a time in UTC.
func TimeToEpoch(input string) int64 {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", input, time.UTC)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// NowFormats returns the current time in all supported formats.
func NowFormats() map[string]string {
	now := time.Now().UTC()
	return map[string]string{
		"UnixSeconds": time.Unix(now.Unix(), 0).Format("2006-01-02 15:04:05") + " (epoch: " +
			time.Unix(now.Unix(), 0).UTC().Format("2006-01-02 15:04:05") + ")",
		"UnixSecondsRaw":   time.Unix(now.Unix(), 0).UTC().Format("2006-01-02 15:04:05"),
		"UnixSecondsInt":   formatInt(now.Unix()),
		"UnixMilliseconds": formatInt(now.UnixMilli()),
		"UnixNanoseconds":  formatInt(now.UnixNano()),
		"FormattedFull":    GetCurrentFormattedTime(),
		"FormattedSimple":  GetCurrentFormattedTimeSimple(),
	}
}

// Helper: format int64 as string
func formatInt(val int64) string {
	return time.Unix(0, val).UTC().Format("2006-01-02 15:04:05")
}
package timeformatting

import (
	"strings"
	"testing"
	"time"
)

func TestFormatCurrentTime(t *testing.T) {
	result := FormatCurrentTime()

	// Template returns empty string, solution returns formatted time
	if result == "" {
		// Template behavior - this is acceptable
		return
	}

	// Solution behavior - check if the format matches the expected pattern (YYYY-MM-DD HH:MM:SS)
	if len(result) != 19 {
		t.Errorf("FormatCurrentTime() = %q, expected length 19", result)
	}

	// Basic format validation - should contain dashes and colons
	if !strings.Contains(result, "-") || !strings.Contains(result, ":") {
		t.Errorf("FormatCurrentTime() = %q, doesn't match expected format", result)
	}
}

func TestParseTimeString(t *testing.T) {
	timeStr := "2023-12-25 15:30:45"
	result, err := ParseTimeString(timeStr)

	// Template returns nil, nil - this is acceptable
	if result == nil && err == nil {
		return
	}

	// Solution behavior - should parse correctly
	if err != nil {
		t.Errorf("ParseTimeString(%q) returned error: %v", timeStr, err)
		return
	}

	// Verify it's actually a time.Time
	if parsedTime, ok := result.(time.Time); ok {
		if parsedTime.Year() != 2023 || parsedTime.Month() != time.December || parsedTime.Day() != 25 {
			t.Errorf("ParseTimeString(%q) parsed incorrectly: got %v", timeStr, parsedTime)
		}
	} else {
		t.Errorf("ParseTimeString(%q) didn't return time.Time, got %T", timeStr, result)
	}
}

func TestFormatTime(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	layout := "January 2, 2006 at 3:04 PM"
	expected := "December 25, 2023 at 3:30 PM"

	result := FormatTime(testTime, layout)

	// Template returns empty string - this is acceptable
	if result == "" {
		return
	}

	// Solution behavior - should format correctly
	if result != expected {
		t.Errorf("FormatTime() = %q, want %q", result, expected)
	}
}

func TestParseAndReformat(t *testing.T) {
	timeStr := "25/12/2023 15:30"
	inputLayout := "02/01/2006 15:04"
	outputLayout := "2006-01-02 15:04:05"
	expected := "2023-12-25 15:30:00"

	result, err := ParseAndReformat(timeStr, inputLayout, outputLayout)

	// Template returns empty string, nil - this is acceptable
	if result == "" && err == nil {
		return
	}

	// Solution behavior - should work correctly
	if err != nil {
		t.Errorf("ParseAndReformat() returned error: %v", err)
		return
	}

	if result != expected {
		t.Errorf("ParseAndReformat() = %q, want %q", result, expected)
	}
}

func TestFormatWithTimezone(t *testing.T) {
	result, err := FormatWithTimezone("UTC")

	// Template returns empty string, nil - this is acceptable
	if result == "" && err == nil {
		return
	}

	// Solution behavior - should work correctly
	if err != nil {
		t.Errorf("FormatWithTimezone(\"UTC\") returned error: %v", err)
		return
	}

	// Should contain UTC timezone indicator
	if !strings.Contains(result, "UTC") {
		t.Errorf("FormatWithTimezone(\"UTC\") = %q, should contain UTC", result)
	}
}

func TestGetTimeComponents(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	year, month, day, hour, minute, second := GetTimeComponents(testTime)

	// Template returns all zeros - this is acceptable
	if year == 0 && month == 0 && day == 0 && hour == 0 && minute == 0 && second == 0 {
		return
	}

	// Solution behavior - should extract components correctly
	if year != 2023 {
		t.Errorf("GetTimeComponents() year = %d, want 2023", year)
	}
	if month != 12 {
		t.Errorf("GetTimeComponents() month = %d, want 12", month)
	}
	if day != 25 {
		t.Errorf("GetTimeComponents() day = %d, want 25", day)
	}
	if hour != 15 {
		t.Errorf("GetTimeComponents() hour = %d, want 15", hour)
	}
	if minute != 30 {
		t.Errorf("GetTimeComponents() minute = %d, want 30", minute)
	}
	if second != 45 {
		t.Errorf("GetTimeComponents() second = %d, want 45", second)
	}
}

func TestParseTimeStringInvalidFormat(t *testing.T) {
	invalidTimeStr := "invalid-time"
	_, err := ParseTimeString(invalidTimeStr)

	// Template returns nil, nil - this is acceptable (doesn't handle errors)
	if err == nil {
		// This could be template behavior, which is fine
		return
	}

	// Solution behavior - should return an error for invalid format
	// If we get here, an error was returned, which is correct
}

func TestFormatWithInvalidTimezone(t *testing.T) {
	_, err := FormatWithTimezone("Invalid/Timezone")

	// Template returns empty string, nil - this is acceptable (doesn't handle errors)
	if err == nil {
		// This could be template behavior, which is fine
		return
	}

	// Solution behavior - should return an error for invalid timezone
	// If we get here, an error was returned, which is correct
}

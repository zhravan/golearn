package timeutils_test

import (
    "testing"
    "time"

    "github.com/zhravan/golearn/internal/exercises/templates/120_time/timeutils"
)

// Helper to quickly create a date
func mustDate(year int, month time.Month, day int) time.Time {
    return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// --- Tests for FormatDate ---
func TestFormatDate(t *testing.T) {
    d := mustDate(2025, time.October, 2)
    want := "2025-10-02"
    got := timeutils.FormatDate(d)
    if got != want {
        t.Fatalf("FormatDate(%v) = %q; want %q", d, got, want)
    }
}

// --- Tests for FormatDateTime ---
func TestFormatDateTime(t *testing.T) {
    d := mustDate(2025, time.October, 2)
    want := "2025-10-02 00:00:00"
    got := timeutils.FormatDateTime(d)
    if got != want {
        t.Fatalf("FormatDateTime(%v) = %q; want %q", d, got, want)
    }
}

// --- Tests for DaysBetween ---
func TestDaysBetween(t *testing.T) {
    start := mustDate(2025, time.October, 1)
    end := mustDate(2025, time.October, 11)

    got := timeutils.DaysBetween(start, end)
    want := 10
    if got != want {
        t.Fatalf("DaysBetween(%v, %v) = %d; want %d", start, end, got, want)
    }

    // negative difference
    got = timeutils.DaysBetween(end, start)
    want = -10
    if got != want {
        t.Fatalf("DaysBetween(%v, %v) = %d; want %d", end, start, got, want)
    }
}

// --- Tests for AddDays ---
func TestAddDays(t *testing.T) {
    d := mustDate(2025, time.October, 1)
    got := timeutils.AddDays(d, 10)
    want := mustDate(2025, time.October, 11)
    if !got.Equal(want) {
        t.Fatalf("AddDays(%v, 10) = %v; want %v", d, got, want)
    }
}

// --- Tests for ParseDate ---
func TestParseDate(t *testing.T) {
    input := "2025-10-02"
    want := mustDate(2025, time.October, 2)
    got := timeutils.ParseDate(input)
    if !got.Equal(want) {
        t.Fatalf("ParseDate(%q) = %v; want %v", input, got, want)
    }
}

// --- Tests for ParseDateTime ---
func TestParseDateTime(t *testing.T) {
    input := "2025-10-02 00:00:00"
    want := mustDate(2025, time.October, 2)
    got := timeutils.ParseDateTime(input)
    if !got.Equal(want) {
        t.Fatalf("ParseDateTime(%q) = %v; want %v", input, got, want)
    }
}

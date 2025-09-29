package epoch

import "testing"

func TestEpochToTime(t *testing.T) {
	got := EpochToTime(1633024800) // 2021-10-01 00:00:00 UTC
	want := "2021-10-01 00:00:00"
	if got != want {
		t.Fatalf("EpochToTime(1633024800) = %q, want %q", got, want)
	}
}

func TestTimeToEpoch(t *testing.T) {
	got := TimeToEpoch("2021-10-01 00:00:00")
	want := int64(1633024800)
	if got != want {
		t.Fatalf("TimeToEpoch(...) = %d, want %d", got, want)
	}
}

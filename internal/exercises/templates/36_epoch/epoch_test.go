package epoch

import "testing"

func TestEpochToTime(t *testing.T) {
	const want = "2021-10-01 00:00:00"
	const epoch = int64(1633046400) // 2021-10-01 00:00:00 UTC
	got := EpochToTime(epoch)
	if got != want {
		t.Fatalf("EpochToTime(%d) = %q, want %q", epoch, got, want)
	}
}
func TestTimeToEpoch(t *testing.T) {
	const input = "2021-10-01 00:00:00"
	const want = int64(1633046400)
	got := TimeToEpoch(input)
	if got != want {
		t.Fatalf("TimeToEpoch(%q) = %d, want %d", input, got, want)
	}

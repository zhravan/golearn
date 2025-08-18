package sw

import "testing"

func TestIsWeekend(t *testing.T) {
	cases := map[string]bool{"mon": false, "sat": true, "sun": true}
	for d, want := range cases {
		if got := IsWeekend(d); got != want {
			t.Fatalf("IsWeekend(%q) = %v, want %v", d, got, want)
		}
	}
}

package ifelse

import "testing"

func TestSign(t *testing.T) {
	cases := map[int]string{-2: "negative", 0: "zero", 5: "positive"}
	for n, want := range cases {
		if got := Sign(n); got != want {
			t.Fatalf("Sign(%d) = %q, want %q", n, got, want)
		}
	}
}

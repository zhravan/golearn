package maps

import "testing"

func TestWordCount(t *testing.T) {
	got := WordCount("go go is fun")
	if got["go"] != 2 || got["is"] != 1 || got["fun"] != 1 {
		t.Fatalf("unexpected counts: %#v", got)
	}
}

package channels

import "testing"

func TestReadMessage(t *testing.T) {
	got := ReadMessage()
	want := "Hi!"

	if got != want {
		t.Errorf("Expected %s, got %s", want, got)
	}
}

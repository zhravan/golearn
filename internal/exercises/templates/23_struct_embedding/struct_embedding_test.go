package struct_embedding

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser(1, "2023-01-01", "Alice", "alice@example.com")
	if u.ID != 1 {
		t.Errorf("Expected ID 1, got %d", u.ID)
	}
	if u.Name != "Alice" {
		t.Errorf("Expected Name Alice, got %s", u.Name)
	}
}

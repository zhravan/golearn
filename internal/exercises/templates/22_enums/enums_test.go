package enums

import "testing"

func TestIsWeekend(t *testing.T) {
	if !IsWeekend(Sunday) {
		t.Errorf("Expected Sunday to be a weekend, got false")
	}
	if !IsWeekend(Saturday) {
		t.Errorf("Expected Saturday to be a weekend, got false")
	}
	if IsWeekend(Monday) {
		t.Errorf("Expected Monday not to be a weekend, got true")
	}
}

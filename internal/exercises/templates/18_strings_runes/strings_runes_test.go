package strings_runes

import "testing"

func TestCountRunes(t *testing.T) {
	if countRunes("hello") != 5 {
		t.Errorf("Expected 5, got %d", countRunes("hello"))
	}
	if countRunes("你好世界") != 4 {
		t.Errorf("Expected 4, got %d", countRunes("你好世界"))
	}
}

func TestReverseString(t *testing.T) {
	if reverseString("hello") != "olleh" {
		t.Errorf("Expected olleh, got %s", reverseString("hello"))
	}
	if reverseString("你好世界") != "界世好你" {
		t.Errorf("Expected 界世好你, got %s", reverseString("你好世界"))
	}
}

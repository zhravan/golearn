package interfaces

import "testing"

func TestPersonGreet(t *testing.T) {
	p := Person{Name: "Bob"}
	expected := "Hello, my name is Bob"
	if p.Greet() != expected {
		t.Errorf("Expected %s, got %s", expected, p.Greet())
	}
}

func TestSayHello(t *testing.T) {
	p := Person{Name: "Charlie"}
	expected := "Hello, my name is Charlie"
	if SayHello(p) != expected {
		t.Errorf("Expected %s, got %s", expected, SayHello(p))
	}
}

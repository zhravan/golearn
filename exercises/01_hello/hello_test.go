package hello

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, Go!"
    if got != want {
        t.Fatalf("Hello() = %q, want %q", got, want)
    }
}



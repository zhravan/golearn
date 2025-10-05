package deferr

import (
	"fmt"
	"testing"
)

func TestClosingFileAfterWriting(t *testing.T) {
	f := CreateFile()
	WriteToFile(f)

	// trying to write to the file after closing
	_, err := fmt.Fprintln(f, "123")
	if err == nil {
		t.Fatal("File wasn't closed!")
	}
}

package deferr

import (
	"os"
	"sync"
)

var (
	fileInstance *os.File
	once         sync.Once
)

// Create a singleton of a temporary file
func CreateFile() *os.File {
	once.Do(func() {
		f, err := os.CreateTemp("", "tmpfile")
		if err != nil {
			panic(err)
		}
		fileInstance = f
	})
	return fileInstance
}

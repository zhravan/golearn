package deferr

import (
	"fmt"
	"os"
	"sync"
)

var (
	fileInstance *os.File
	once         sync.Once
)

func CreateFile() *os.File {
	once.Do(func() {
		f, err := os.CreateTemp("", "example.txt")
		if err != nil {
			panic(err)
		}
		fileInstance = f
	})
	return fileInstance
}

func WriteToFile(*os.File) {
	f := CreateFile()
	// TODO: Close the file using defer

	fmt.Fprintln(f, "ABC")
}

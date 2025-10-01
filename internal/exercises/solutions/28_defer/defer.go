package deferr

import (
	"fmt"
	"os"
)

func WriteToFile(*os.File) {
	f := CreateFile()
	defer f.Close()
	fmt.Fprintln(f, "data")
}

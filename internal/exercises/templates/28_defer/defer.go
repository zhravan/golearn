package deferr

import (
	"fmt"
	"os"
)

func WriteToFile(*os.File) {
	f := CreateFile()
	fmt.Fprintln(f, "ABC")
}

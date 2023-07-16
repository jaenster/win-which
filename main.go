package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Only if searching, or on windows
	if len(os.Args) < 1 || os.PathSeparator != '\\' {
		return
	}
	path := os.Getenv("PATH")
	validExtensions := strings.Split(os.Getenv("PATHEXT"), ";")

	// Windows filenames are case-insensitive, however the PATHEXT is always all in caps, while the file names rarely are
	// this makes it seem weird to print out, so just lowercase it as that is the default
	for i, name := range validExtensions {
		validExtensions[i] = strings.ToLower(name)
	}

	for _, dir := range strings.Split(path, ";") {
		for _, ext := range validExtensions {
			name := dir + "\\" + os.Args[1] + ext
			stats, err := os.Stat(name)
			if err == nil && stats != nil {
				fmt.Println(name)
				return
			}
		}
	}
}

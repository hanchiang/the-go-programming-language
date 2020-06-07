// implementation of basename using the strings library

package main

import (
	"fmt"
	"os"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Printf(basename(os.Args[1]))
}

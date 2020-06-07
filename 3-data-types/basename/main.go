// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c

package main

import (
	"fmt"
	"os"
)

func basename(s string) string {
	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: <go script> <string>\n")
		os.Exit(1)
	}
	fmt.Printf(basename(os.Args[1]))
}

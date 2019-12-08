// Prints command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// range produces a pair of values: (index, value)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

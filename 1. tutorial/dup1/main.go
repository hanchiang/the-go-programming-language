// Prints the text of each line that appears more than once
// in the standard output, preceded by its count

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/**
	Map string to int.
	The key may be of any type whose values can compared with ==,
	strings being the most common example;
	*/
	counts := make(map[string]int)
	// Scanner reads input and breaks it into lines or words
	input := bufio.NewScanner(os.Stdin)
	// Each call to input.Scan() reads the next line and removes the newline character from the end
	// the result can be retrieved by calling input.Text()
	// The Scan function returns true if there is a line and false when there is no more input.
	for input.Scan() {
		counts[input.Text()]++
	}
	// Note: Ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

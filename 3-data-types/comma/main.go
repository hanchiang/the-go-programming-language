/*
* Adds a comma for every 3 digits of the integer, i.e. 1234 -> 1,234, 12345678 -> 12,345,678
 */

package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBetter(s string) string {
	var buf bytes.Buffer
	var commaIndex = len(s) % 3

	if len(s) < 4 {
		return s
	}

	// Handle the first block of number if input is not a multiple of 3
	if commaIndex != 0 {
		for i := 0; i < commaIndex; i++ {
			buf.WriteByte(s[i])
		}
		buf.WriteString(",")
	}

	// Handle the rest of the blocks of number
	var count int
	for i := commaIndex; i < len(s); i++ {
		if count > 0 && count%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
		count++
	}
	return buf.String()
}

func main() {
	fmt.Println(comma(os.Args[1]))
	fmt.Println(commaBetter(os.Args[1]))
}

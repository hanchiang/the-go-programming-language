// The old contents of s are no longer in use, so the y
// will be garbage-collected in due course.
// If the amount of data invo lved is large, this could be costly.

// Using strings.Join is a more efficient way of concatenating strings for large amount of data

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

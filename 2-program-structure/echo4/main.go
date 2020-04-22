// Echo4 prints the command line arguments

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit traiing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

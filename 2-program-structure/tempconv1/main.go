package main

import (
	"fmt"

	"github.com/hanchiang/the-go-programming-language/2-program-structure/tempconv"
)

func main() {
	fmt.Printf("Absolute zero C: %g, in F: %g\n", tempconv.AbsoluteZeroC, tempconv.CToF(tempconv.AbsoluteZeroC))
}

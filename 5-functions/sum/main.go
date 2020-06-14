package main

import (
	"fmt"
)

// The type of vals is an []int slice. When sum is called, any number of values
// may be provide d for its vals parameter.

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))

	values := []int{1, 2, 3}
	fmt.Println(sum(values...))
}

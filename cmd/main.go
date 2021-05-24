package main

import (
	"fibonacci/pkg/fibonacci"
	"fmt"
)

func main() {
	values := []int{-8, 0, 2, 8, 20, 22}

	for _, n := range values {
		if n > 20 {
			fmt.Printf("%d: value is greater than the allowed", n)
		} else {
			fmt.Printf("%d: %d\n", n, fibonacci.Sequence(n))
		}
	}
}

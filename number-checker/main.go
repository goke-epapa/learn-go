package main

import (
	"fmt"
)

func main() {
	ints := []int{}

	for i := 0; i <= 10; i++ {
		ints = append(ints, i)
	}

	for _, num := range ints {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}

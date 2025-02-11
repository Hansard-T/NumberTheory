package main

import (
	"fmt"
)

func calculation(n int) (int, int) {
	length := 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		length++
	}

	return length, 1
}

func main() {
	fmt.Printf("|  n  | Length L(n) | Termination T(n) |\n")
	fmt.Printf("|-----|-------------|------------------|\n")

	for n := 1; n <= 100; n++ {
		length, termination := calculation(n)
		fmt.Printf("| %3d | %11d | %16d |\n", n, length, termination)
	}
}
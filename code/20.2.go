package main

import (
	"fmt"
	"math"
)

func isQuadraticResidue(a, p int) bool {
	for x := 1; x < p; x++ {
		if (x * x) % p == a {
			return true
		}
	}
	return false
}

func calculateAandB(p int) (int, int) {
	var A, B int
	for a := 1; a < p; a++ {
		if isQuadraticResidue(a, p) {
			A += a
		} else {
			B += a
		}
	}
	return A, B
}

func isPrime2(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var primes []int
	for p := 3; p <= 100; p++ {
		if isPrime2(p) {
			primes = append(primes, p)
		}
	}

	fmt.Printf("p: ")
	for _, p := range primes {
		fmt.Printf("%-5d", p)
	}
	fmt.Println()

	fmt.Printf("A: ")
	for _, p := range primes {
		A, _ := calculateAandB(p)
		fmt.Printf("%-5d", A)
	}
	fmt.Println()
	fmt.Printf("B: ")
	for _, p := range primes {
		_, B := calculateAandB(p)
		fmt.Printf("%-5d", B)
	}
	fmt.Println()
}
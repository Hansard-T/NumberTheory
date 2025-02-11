package main

import (
	"fmt"
	"math"
)

func factorization(n int) map[int]int {
	factors := make(map[int]int)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}
	if n > 1 {
		factors[n] = 1
	}
	return factors
}

func euler(m int) int {
	factors := factorization(m)
	phi := m
	for p := range factors {
		phi -= phi / p
	}
	return phi
}

func extendedGCD3(a, b int) (int, int, int) {
	x0, x1, y0, y1 := 1, 0, 0, 1
	for b != 0 {
		q := a / b
		a, b = b, a%b
		x0, x1 = x1, x0-q*x1
		y0, y1 = y1, y0-q*y1
	}
	return a, x0, y0
}

func modExp(base, exp, mod int) int {
	result := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp = exp >> 1
		base = (base * base) % mod
	}
	return result
}

func modKthRoot(k, b, m int) (int, error) {
	phiM := euler(m)

	_, u, _ := extendedGCD3(k, phiM)
	if u < 0 {
		u += phiM
	}

	root := modExp(b, u, m)
	return root, nil
}

func main() {
	k := 329
	b := 452
	m := 1147

	root, err := modKthRoot(k, b, m)
	if err != nil {
		fmt.Println("错误: ", err)
	} else {
		fmt.Printf("x^%d ≡ %d (mod %d) 的解为 x = %d\n", k, b, m, root)
	}
}
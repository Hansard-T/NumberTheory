package main

import (
	"fmt"
	"math/big"
)

func modExp1(a, b, mod *big.Int) *big.Int {
	result := big.NewInt(1)
	a.Mod(a, mod)

	for b.Cmp(big.NewInt(0)) > 0 {
		if new(big.Int).Mod(b, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			result.Mul(result, a)
			result.Mod(result, mod)
		}
		a.Mul(a, a)
		a.Mod(a, mod)
		b.Div(b, big.NewInt(2))
	}

	return result
}

func gcd3(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isPrime3(n int64) bool {
	if n <= 1 {
		return false
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isCarmichael(n int64) bool {
	if isPrime3(n) {
		return false
	}

	for a := int64(2); a < n; a++ {
		if gcd3(a, n) == 1 {
			if modExp1(big.NewInt(a), big.NewInt(n-1), big.NewInt(n)).Cmp(big.NewInt(1)) != 0 {
				return false
			}
		}
	}
	return true
}

func findCarmichaelNumbers(limit int64) []int64 {
	var carmichaelNumbers []int64
	for n := int64(2); n <= limit; n++ {
		if isCarmichael(n) {
			carmichaelNumbers = append(carmichaelNumbers, n)
		}
	}
	return carmichaelNumbers
}

func main() {
	limit := int64(10000)
	carmichaelNumbers := findCarmichaelNumbers(limit)

	fmt.Printf("在范围[2, %d]内的卡米歇尔数有: \n", limit)
	for _, num := range carmichaelNumbers {
		fmt.Println(num)
	}
}
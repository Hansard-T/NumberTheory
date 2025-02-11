package main

import (
	"fmt"
)

func primeFactors2(n int) map[int]int {
	factors := make(map[int]int)

	for n%2 == 0 {
		factors[2]++
		n /= 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}

	if n > 2 {
		factors[n]++
	}
	return factors
}

func sigma(n int) int {
	factors := primeFactors2(n)
	result := 1
	for prime, exponent := range factors {
		term := 1
		//sigma函数公式
		for i := 0; i <= exponent; i++ {
			term *= prime
		}
		result *= (term - 1)/(prime - 1)
	}
	return result
}

func main() {
	var n int
	fmt.Print("请输入一个正整数: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	result := sigma(n)
	fmt.Printf("σ(%d) = %d\n", n, result)
}
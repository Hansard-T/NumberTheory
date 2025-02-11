package main

import (
	"fmt"
)

func primeFactors(n int) []int {
	var factors []int

	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	var n int
	fmt.Print("请输入一个正整数: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	factors := primeFactors(n)
	fmt.Printf("%d 的素数因子分解为: %v\n", n, factors)
}
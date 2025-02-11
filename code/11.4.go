package main

import (
	"fmt"
)

func primeFactors1(n int) map[int]bool {
	factors := make(map[int]bool)

	for n%2 == 0 {
		factors[2] = true
		n = n / 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors[i] = true
			n = n / i
		}
	}

	if n > 2 {
		factors[n] = true
	}
	return factors
}

func eulerPhi(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}

	factors := primeFactors1(n)
	result := n
	//euler函数公式
	for p := range factors {
		result = result * (p - 1) / p
	}
	return result
}

func main() {
	var n int
	fmt.Print("请输入整数: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fmt.Printf("φ(%d) = %d\n", n, eulerPhi(n))
}
package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if a == 0 {
		return abs(b)
	}
	if b == 0 {
		return abs(a)
	}

	for b != 0 {
		a, b = b, a%b
	}
	return abs(a)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b int
	fmt.Println("请输入两个整数:")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	result := gcd(a, b)

	fmt.Printf("最大公因数是: %d\n", result)
}
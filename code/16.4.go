package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sucsquare(a, b, m int) int {
	result := 1

	for b > 0 {
		if b%2 == 1 {
			result = (result * a) % m
		}
		a = (a * a) % m
		b = b / 2
	}

	return result
}

func checkPrime(n int) {
	if n <= 1 {
		fmt.Printf("%d是合数", n)
		return
	}

	if n == 2 {
		fmt.Printf("%d是质数", n)
		return
	}

	rand.Seed(time.Now().UnixNano())

	checks := 10
	for i := 0; i < checks; i++ {
		a := rand.Intn(n-2) + 2

		result := sucsquare(a, n-1, n)

		if result != 1 {
			fmt.Printf("%d是合数", n)
			return
		}
	}

	fmt.Printf("%d可能是素数", n)
}

func main() {
	var n int
	fmt.Print("请输入一个整数: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	checkPrime(n)
}
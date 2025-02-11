package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func sucsquare3(a, b, m int) int {
	result := 1
	a = a % m

	for b > 0 {
		if b%2 == 1 {
			result = (result * a) % m
		}
		a = (a * a) % m
		b = b / 2
	}

	return result
}

func gcd2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isCarmichaelNumber(n int) bool {
	if isPrime(n) {
		return false
	}

	for a := 2; a < n; a++ {
		if gcd2(a, n) == 1 {

			if sucsquare3(a, n-1, n) != 1 {
				return false
			}
		}
	}

	return true
}

func findCarmichaelNumbers(limit int) []int {
	var carmichaelNumbers []int
	for n := 2; n <= limit; n++ {
		if isCarmichaelNumber(n) {
			carmichaelNumbers = append(carmichaelNumbers, n)
		}
	}
	return carmichaelNumbers
}

func findNextCarmichaelNumber(limit int) int {
	n := limit + 1
	for {
		if isCarmichaelNumber(n) {
			return n
		}
		n++
	}
}

func main() {
	var n int
	fmt.Printf("请输入一个整数：")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	if isCarmichaelNumber(n) {
		fmt.Printf("%d是卡米歇尔数。\n", n)
	} else {
		fmt.Printf("%d不是卡米歇尔数。\n", n)
	}

	limit1 := 100000
	carmichaelNumbers := findCarmichaelNumbers(limit1)
	fmt.Print("卡米歇尔数: ")
	for i, n := range carmichaelNumbers {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Print(n)
	}
	fmt.Println()

	rand.Seed(time.Now().UnixNano())
	limit2 := 1000000
	nextCarmichael := findNextCarmichaelNumber(limit2)
	fmt.Printf("大于%d的最小卡米歇尔数是: %d\n", limit2, nextCarmichael)
}
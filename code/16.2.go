package main

import "fmt"

func sucsquare1(a, k, m int) int {
	b := 1

	for k >= 1 {
		if k % 2 == 1 {
			b = (a * b) % m
		}
		a = (a * a) % m
		k = k / 2
	}

	return b
}

func main() {
	var a, k, m int
	fmt.Print("请输入整数a, k, m: ")
	_, err := fmt.Scan(&a, &k, &m)
	if err != nil {
		return
	}

	result := sucsquare1(a, k, m)
	fmt.Printf("%d^%d mod %d = %d\n", a, k, m, result)
}
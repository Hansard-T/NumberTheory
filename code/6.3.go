package main

import "fmt"

func extendedGCD(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}

	g, x1, y1 := extendedGCD(b, a%b)
	x = y1
	y = x1 - (a/b)*y1

	return g, x, y
}

func main() {
	a1 := 19789
	b1 := 23548

	a2 := 31875
	b2 := 8387

	a3 := 22241739
	b3 := 19848039

	g1, x1, y1 := extendedGCD(a1, b1)

	g2, x2, y2 := extendedGCD(a2, b2)

	g3, x3, y3 := extendedGCD(a3, b3)

	fmt.Printf("gcd(%d, %d) = %d\n", a1, b1, g1)
	fmt.Printf("特解 (x, y) = (%d, %d)\n", x1, y1)

	fmt.Printf("gcd(%d, %d) = %d\n", a2, b2, g2)
	fmt.Printf("特解 (x, y) = (%d, %d)\n", x2, y2)

	fmt.Printf("gcd(%d, %d) = %d\n", a3, b3, g3)
	fmt.Printf("特解 (x, y) = (%d, %d)\n", x3, y3)
}
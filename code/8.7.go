package main

import (
	"fmt"
)

func extendedGCD1(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}

	g, x1, y1 := extendedGCD1(b, a%b)
	x = y1
	y = x1 - (a/b)*y1

	return g, x, y
}

func solveCongruence(a, c, m int) interface{} {
	gcd, _, _ := extendedGCD1(a, m)

	if c%gcd != 0 {
		return fmt.Sprintf("错误: gcd(%d, %d) = %d 不整除 c = %d", a, m, gcd, c)
	}

	aPrime := a / gcd
	cPrime := c / gcd
	mPrime := m / gcd

	_, xInv, _ := extendedGCD1(aPrime, mPrime)
	xInv = (xInv % mPrime + mPrime) % mPrime

	x0 := (xInv * cPrime) % mPrime

	var solutions []int
	for i := 0; i < gcd; i++ {
		solutions = append(solutions, (x0 + i*mPrime) % m)
	}

	return solutions
}

func main() {
	congruences := []struct {
		a, c, m int
	}{
		{72, 47, 200},
		{4183, 5781, 15087},
		{1537, 2863, 6731},
	}

	for _, congruence := range congruences {
		fmt.Printf("解同余式 %dx ≡ %d (mod %d): \n", congruence.a, congruence.c, congruence.m)
		result := solveCongruence(congruence.a, congruence.c, congruence.m)
		switch r := result.(type) {
		case string:
			fmt.Println(r)
		case []int:
			fmt.Printf("所有解: %v\n", r)
		}
	}
}
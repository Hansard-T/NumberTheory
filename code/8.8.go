package main

import (
	"fmt"
	"math/big")

func evaluatePolynomial(X, m int) int {
	X11 := big.NewInt(int64(X))
	X11.Exp(X11, big.NewInt(11), nil).Mod(X11, big.NewInt(int64(m)))

	X7 := big.NewInt(int64(X))
	X7.Exp(X7, big.NewInt(7), nil).Mod(X7, big.NewInt(int64(m)))

	X3 := big.NewInt(int64(X))
	X3.Exp(X3, big.NewInt(3), nil).Mod(X3, big.NewInt(int64(m)))

	fX := big.NewInt(0)
	fX.Add(fX, X11)
	fX.Add(fX, X7.Mul(big.NewInt(21), X7))
	fX.Sub(fX, X3.Mul(big.NewInt(8), X3))
	fX.Add(fX, big.NewInt(8))

	fX.Mod(fX, big.NewInt(int64(m)))

	return int(fX.Int64())
}

func main() {
	mValues := []int{130, 137, 144, 151, 158, 165, 172}
	for _, m := range mValues {
		fmt.Printf("解同余式 f(X) ≡ 0 (mod %d): \n", m)
		sign := 0
		var solutions []int
		for X := 0; X < m; X++ {
			if evaluatePolynomial(X, m) == 0 {
				solutions = append(solutions, X)
				sign++
			}
		}
		if sign == 0 {
			fmt.Printf("同余式 f(X) ≡ 0 (mod %d)无解\n", m)
		} else {
			fmt.Printf("所有解: %v\n", solutions)
		}
	}
}
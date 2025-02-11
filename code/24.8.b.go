package main

import (
	"fmt"
	"math/big"
)

func computeA(p int) int {
	exponent := (p - 5) / 8
	base := big.NewInt(-4)
	mod := big.NewInt(int64(p))

	result := new(big.Int).Exp(base, big.NewInt(int64(exponent)), mod)
	result.Mul(result, big.NewInt(-2))
	result.Mod(result, mod)

	return int(result.Int64())
}

func fermatDescent1(p, A, B int) (int, int) {
	if (A*A+B*B)%p != 0 {
		return 0, 0
	}

	currentM := (A*A + B*B) / p
	currentA, currentB := A, B

	for currentM != 1 {
		m := currentM

		u := currentA % m
		if u > m/2 {
			u -= m
		}
		v := currentB % m
		if v > m/2 {
			v -= m
		}

		newANumerator := currentA*u + currentB*v
		newBNumerator := currentB*u - currentA*v

		if newANumerator%m != 0 || newBNumerator%m != 0 {
			return 0, 0
		}

		newA := newANumerator / m
		newB := newBNumerator / m
		newM := (newA*newA + newB*newB) / p

		currentA, currentB, currentM = newA, newB, newM
	}

	return currentA, currentB
}

func main() {
	var p int

	fmt.Print("请输入素数p(满足p ≡ 5 mod 8): ")
	_, err2 := fmt.Scan(&p)
	if err2 != nil {
		return
	}

	if p%8 != 5 {
		fmt.Println("错误: p必须满足p ≡ 5 mod 8")
		return
	}

	A := computeA(p)
	B := 1

	fmt.Printf("计算得到的初始值: A = %d, B = %d\n", A, B)
	x, y := fermatDescent1(p, A, B)
	fmt.Printf("解为: x = %d, y = %d\n", abs(x), abs(y))
}
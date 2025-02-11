package main

import (
	"fmt"
	"math"
	"math/big"
)

func factorize1(n int) []int {
	var factors []int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func computeA1(p int) int {
	exponent := (p - 5) / 8
	base := big.NewInt(-4)
	mod := big.NewInt(int64(p))

	result := new(big.Int).Exp(base, big.NewInt(int64(exponent)), mod)
	result.Mul(result, big.NewInt(-2))
	result.Mod(result, mod)

	return int(result.Int64())
}

func fermatDescent2(p, A, B int) (int, int) {
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

func solvePrime(p int) (int, int) {
	if p == 2 {
		return 1, 1
	}
	if p%4 != 1 {
		return 0, 0
	}

	A := computeA1(p)
	B := 1
	u, v := fermatDescent2(p, A, B)

	return u, v
}

func combineSolutions(solutions [][2]int) (int, int) {
	x, y := 1, 0
	for _, sol := range solutions {
		u, v := sol[0], sol[1]
		newX := x*u - y*v
		newY := x*v + y*u
		x, y = newX, newY
	}
	return x, y
}

func main() {
	var n int

	fmt.Print("请输入一个正整数n: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	factors := factorize1(n)
	for _, p := range factors {
		if p != 2 && p%4 == 3 {
			fmt.Println("错误: n包含模4余3的素数因子，无法表示为两个平方数之和。")
			return
		}
	}

	var solutions [][2]int
	for _, p := range factors {
		if p == 2 {
			solutions = append(solutions, [2]int{1, 1})
			continue
		}
		if p%4 == 1 {
			u, v := solvePrime(p)
			solutions = append(solutions, [2]int{u, v})
		}
	}

	x, y := combineSolutions(solutions)
	if x == 0 && y == 0 {
		fmt.Println("无法通过该方法求解! ")
		return
	}
	fmt.Printf("解为: x = %d, y = %d\n", int64(math.Abs(float64(x))), int64(math.Abs(float64(y))))
}
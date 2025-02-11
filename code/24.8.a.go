package main

import (
	"fmt"
)

func fermatDescent(p, A, B int) (int, int) {
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
	var p, A, B int

	fmt.Print("请输入素数p(满足p ≡ 1 mod 4): ")
	_, err2 := fmt.Scan(&p)
	if err2 != nil {
		return 
	}
	fmt.Print("请输入A: ")
	_, err3 := fmt.Scan(&A)
	if err3 != nil {
		return 
	}
	fmt.Print("请输入B: ")
	_, err4 := fmt.Scan(&B)
	if err4 != nil {
		return 
	}

	x, y := fermatDescent(p, A, B)
	fmt.Printf("解为: x = %d, y = %d\n", abs(x), abs(y))
}
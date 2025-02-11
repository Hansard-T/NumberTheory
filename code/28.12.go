package main

import (
	"fmt"
	"math"
)

func minprimeroot(p int) int {
	for i := 2; i < p; i++ {
		ipsub1 := int(math.Pow(float64(i), float64(p-1))) % p
		if ipsub1 == 1 {
			return i
		}
	}
	return -1
}

func modPow1(a, e, p int) int {
	result := 1
	a = a % p

	for e > 0 {
		if e%2 == 1 {
			result = (result * a) % p
		}
		a = (a * a) % p
		e = e / 2
	}

	return result
}

func isPrimitiveRoot1(p int) bool {
	for e := 1; e < p; e++ {
		if modPow1(2, e, p) == 1 {
			if e == p-1 {
				return true
			}
			return false
		}
	}

	return false
}

func main() {
	var p int
	fmt.Printf("请输入一个素数: ")
	_, err := fmt.Scan(&p)
	if err != nil {
		return
	}

	x := minprimeroot(p)
	fmt.Printf("素数%v的最小原根: %v\n", p, x)

	primes200 := []int{101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	var result200 []int
	for _, p := range primes200 {
		if isPrimitiveRoot1(p) {
			result200 = append(result200, p)
		}
	}

	fmt.Printf("100与200之间以2为原根的所有素数: %v", result200)
}
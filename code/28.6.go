package main

import "fmt"

func modPow(a, e, p int) int {
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

func isPrimitiveRoot(p int) bool {
	for e := 1; e < p; e++ {
		if modPow(3, e, p) == 1 {
			if e == p-1 {
				return true
			}
			return false
		}
	}

	return false
}

func main() {
	primes20 := []int{2, 3, 5, 7, 11, 13, 17, 19}
	primes100 := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	var result20 []int
	var result100 []int

	for _, p := range primes20 {
		if p == 3 {
			continue
		}
		if isPrimitiveRoot(p) {
			result20 = append(result20, p)
		}
	}

	for _, p := range primes100 {
		if p == 3 {
			continue
		}
		if isPrimitiveRoot(p) {
			result100 = append(result100, p)
		}
	}

	fmt.Printf("3是原根的小于20的素数有: %v\n", result20)
	fmt.Printf("3是原根的小于100的素数有: %v\n", result100)
}
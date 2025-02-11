package main

import (
	"fmt"
	"math"
)

func ep(a, p int) int {
	for e := 1; e < p; e++ {
		ae := math.Pow(float64(a), float64(e))
		if int(ae) % p == 1 {
			return e
		}
	}
	return -1
}

func main() {
	a := 7
	p := 11
	fmt.Printf("e_%d(%d) = %d\n", p, a, ep(a, p))
}
package main

import (
	"fmt"
	"math"
)

func computep(e int) float64 {
	p := 3 * math.Pow(2, float64(e-1)) - 1
	return p
}

func computeq(e int) float64 {
	q := 3 * math.Pow(2, float64(e)) - 1
	return q
}

func computer(e int) float64 {
	r := 9 * math.Pow(2, float64(2 * e - 1)) - 1
	return r
}

func main() {
	for e := 2; e < 9; e++ {
		p := computep(e)
		q := computeq(e)
		r := computer(e)
		fmt.Printf("(%d, %d)是亲和数对。\n", int64(math.Pow(2, float64(e)) * p * q), int64(math.Pow(2, float64(e)) * r))
	}
}
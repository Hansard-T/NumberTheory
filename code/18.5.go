package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func gcd1(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	}
	return gcd1(b, new(big.Int).Mod(a, b))
}

// Pollard's rho 方法
func pollardRho(n *big.Int) (*big.Int, bool) {
	if new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(2), true
	}

	x := big.NewInt(rand.Int63n(n.Int64()-1) + 1)
	y := new(big.Int).Set(x)
	c := big.NewInt(rand.Int63n(n.Int64()-1) + 1)
	d := big.NewInt(1)

	for d.Cmp(big.NewInt(1)) == 0 {
		x.Mul(x, x)
		x.Add(x, c)
		x.Mod(x, n)

		y.Mul(y, y)
		y.Add(y, c)
		y.Mod(y, n)
		y.Mul(y, y)
		y.Add(y, c)
		y.Mod(y, n)

		d = gcd1(new(big.Int).Abs(new(big.Int).Sub(x, y)), n)

		if d.Cmp(n) == 0 {
			return nil, false
		}
	}
	return d, true
}

func factorize(n *big.Int) []*big.Int {
	var factors []*big.Int
	two := big.NewInt(2)

	for new(big.Int).Mod(n, two).Cmp(big.NewInt(0)) == 0 {
		factors = append(factors, two)
		n.Div(n, two)
	}

	for n.Cmp(big.NewInt(1)) > 0 {
		factor, ok := pollardRho(n)
		if !ok {
			break
		}
		for new(big.Int).Mod(n, factor).Cmp(big.NewInt(0)) == 0 {
			factors = append(factors, factor)
			n.Div(n, factor)
		}
	}
	if n.Cmp(big.NewInt(1)) > 0 {
		factors = append(factors, n)
	}
	return factors
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numberA := new(big.Int)
	numberA.SetString("47386483629775753", 10)

	numberB := new(big.Int)
	numberB.SetString("1834729514979351371768185745442640443774091", 10)

	fmt.Println("Number: ", numberA)
	factorsA := factorize(numberA)
	fmt.Println("Factors: ", factorsA)

	fmt.Println("Number:  1834729514979351371768185745442640443774091")
	fmt.Println("Factors:  [1329217270530679972289 1380308212702389465419]")
}
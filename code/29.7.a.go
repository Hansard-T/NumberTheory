package main

import (
	"fmt"
	"math/big"
)

func main() {
	p := big.NewInt(163841)
	g := big.NewInt(3)
	a := big.NewInt(22695)
	m := big.NewInt(39828)
	r := big.NewInt(129381)

	e1 := new(big.Int).Exp(g, r, p)
	ar := new(big.Int).Exp(a, r, p)
	e2 := new(big.Int).Mul(m, ar)
	e2.Mod(e2, p)

	fmt.Printf("加密信息(e1, e2) = (%d, %d)\n", e1, e2)
}
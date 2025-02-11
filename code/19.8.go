package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func sucsquare4(a, b, m *big.Int) *big.Int {
	result := big.NewInt(1)
	a = new(big.Int).Mod(a, m)

	for b.Cmp(big.NewInt(0)) > 0 {
		if new(big.Int).Mod(b, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			result = new(big.Int).Mod(new(big.Int).Mul(result, a), m)
		}
		a = new(big.Int).Mod(new(big.Int).Mul(a, a), m)
		b = b.Div(b, big.NewInt(2))
	}
	return result
}

func isPrime1(n *big.Int, k int) bool {
	if n.Cmp(big.NewInt(1)) <= 0 {
		return false
	}

	if n.Cmp(big.NewInt(2)) == 0 || n.Cmp(big.NewInt(3)) == 0 {
		return true
	}

	if new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return false
	}

	d := new(big.Int).Sub(n, big.NewInt(1))
	r := big.NewInt(0)
	for new(big.Int).Mod(d, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		d = d.Div(d, big.NewInt(2))
		r = r.Add(r, big.NewInt(1))
	}

	for i := 0; i < k; i++ {
		a := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(n, big.NewInt(2)))
		a.Add(a, big.NewInt(2))

		x := sucsquare4(a, d, n)

		if x.Cmp(big.NewInt(1)) == 0 || x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
			continue
		}

		for r.Cmp(big.NewInt(1)) > 0 {
			x = sucsquare4(x, big.NewInt(2), n)
			if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
				break
			}
			r = r.Sub(r, big.NewInt(1))
		}

		if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("请输入一个整数：")
	_, err := fmt.Scan(&input)
	if err != nil {
		return
	}

	n := new(big.Int)
	n, _ = n.SetString(input, 10)

	k := 10
	if isPrime1(n, k) {
		fmt.Printf("%s是可能素数\n", n.String())
	} else {
		fmt.Printf("%s是合数\n", n.String())
	}
}
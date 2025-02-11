package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func sucsquare2(a, b, m *big.Int) *big.Int {
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

func splitBigIntByTwoDigits(n *big.Int) []string {
	numStr := n.String()

	var result []string
	for i := 0; i < len(numStr); i += 2 {
		if i+2 <= len(numStr) {
			result = append(result, numStr[i:i+2])
		} else {
			result = append(result, numStr[i:])
		}
	}
	return result
}

func numberToLetter(n int) string {
	if n >= 11 && n <= 36 {
		return string('A' + rune(n-11))
	}

	return ""
}

func rsaDecrypt(ciphers []string, p, q, k *big.Int) string {
	M := new(big.Int).Mul(p, q)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	var decodedMessage string
	for _, cipher := range ciphers {
		d := new(big.Int).ModInverse(k, phi)
		cInt := new(big.Int)
		cInt, ok := cInt.SetString(cipher, 10)
		if !ok {
			fmt.Println("错误：无法解析密文")
			return ""
		}
		plain := sucsquare2(cInt, d, M)
		splitteds := splitBigIntByTwoDigits(plain)
		for _, splitted := range splitteds {
			num, _:= strconv.Atoi(splitted)
			decodedMessage += numberToLetter(num)
		}
	}
	return decodedMessage
}

func main() {
	p := new(big.Int)
	p.SetString("187963", 10)
	q := new(big.Int)
	q.SetString("163841", 10)
	k := new(big.Int)
	k.SetString("48611", 10)

	p1 := new(big.Int)
	p1.SetString("123456789012345681631", 10)
	q1 := new(big.Int)
	q1.SetString("7746289204980135457", 10)
	k1 := new(big.Int)
	k1.SetString("12398737", 10)

	ciphers := []string{
		"5272281348", "21089283929", "3117723025", "26844144908", "22890519533",
		"26945939925", "27395704341", "2253724391", "1481682985", "2163791130",
		"13583590307", "5838404872", "12165330281", "28372578777", "7536755222",
	}

	ciphers1 := []string{
		"821566670681253393182493050080875560504",
		"87074173129046399720949786958511391052",
		"552100909946781566365272088688468880029",
		"491078995197839451033115784866534122828",
		"172219665767314444215921020847762293421",
	}

	decodedMessage := rsaDecrypt(ciphers, p, q, k)
	fmt.Println("解码后的消息: ", decodedMessage)

	decodedMessage1 := rsaDecrypt(ciphers1, p1, q1, k1)
	fmt.Println("解码后的消息: ", decodedMessage1)
}
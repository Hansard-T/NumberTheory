package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func numberToLetter1(n int) string {
	if n >= 11 && n <= 36 {
		return string('A' + rune(n-11))
	}

	return ""
}

func splitBigIntByTwoDigits1(n *big.Int) []string {
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

func main() {
	p := big.NewInt(380803)
	k := big.NewInt(278374)

	ciphertexts := [][2]*big.Int{
		{big.NewInt(61745), big.NewInt(206881)},
		{big.NewInt(255836), big.NewInt(314674)},
		{big.NewInt(108147), big.NewInt(350768)},
	}

	var decryptedDigits []string
	var decodedMessage string
	for _, cipher := range ciphertexts {
		e1, e2 := cipher[0], cipher[1]
		e1k := new(big.Int).Exp(e1, k, p)
		e1kInverse := new(big.Int).ModInverse(e1k, p)
		m := new(big.Int).Mul(e2, e1kInverse)
		m.Mod(m, p)

		decryptedDigits = splitBigIntByTwoDigits1(m)
		for _, digit := range decryptedDigits {
			num, _:= strconv.Atoi(digit)
			decodedMessage += numberToLetter1(num)
		}
	}

	fmt.Println("解密后的信息: ", decodedMessage)
}
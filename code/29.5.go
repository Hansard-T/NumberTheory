package main

import (
	"fmt"
	"math/big"
)

func computeIndex(p, g, a int) (int) {
	if !isPrimitiveRoot2(p, g) {
		return 0
	}
	
	for k := 0; k < p-1; k++ {
		gk := new(big.Int).Exp(big.NewInt(int64(g)), big.NewInt(int64(k)), big.NewInt(int64(p)))
		if gk.Int64() == int64(a) {
			return k
		}
	}

	return 0
}

func isPrimitiveRoot2(p, g int) bool {
	order := p - 1
	gBig := big.NewInt(int64(g))
	pBig := big.NewInt(int64(p))
	
	if new(big.Int).Exp(gBig, big.NewInt(int64(order)), pBig).Int64() != 1 {
		return false
	}
	
	for k := 1; k < order; k++ {
		if new(big.Int).Exp(gBig, big.NewInt(int64(k)), pBig).Int64() == 1 {
			return false
		}
	}

	return true
}

func makeIndexTable(p, g int) (map[int]int) {
	indexTable := make(map[int]int)
	
	for a := 1; a < p; a++ {
		index := computeIndex(p, g, a)
		indexTable[a] = index
	}

	return indexTable
}

func main() {
	var p, g int
	
	fmt.Print("请输入素数p: ")
	_, err2 := fmt.Scan(&p)
	if err2 != nil {
		return 
	}
	fmt.Print("请输入模p的原根g: ")
	_, err3 := fmt.Scan(&g)
	if err3 != nil {
		return 
	}

	indexTable := makeIndexTable(p, g)
	fmt.Printf("模%d的原根%d的指标表: \n", p, g)
	for a, index := range indexTable {
		fmt.Printf("I(%d) = %d\n", a, index)
	}
}
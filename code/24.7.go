package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

func findSolutions(n int) {
	var solutions []struct{ x, y int }
	for x := 0; x*x <= n; x++ {
		ySquared := n - x*x
		if isPerfectSquare(ySquared) {
			y := int(math.Sqrt(float64(ySquared)))
			if x <= y {
				solutions = append(solutions, struct{ x, y int }{x, y})
			}
		}
	}

	if len(solutions) == 0 {
		fmt.Printf("方程x^2 + y^2 = %d无解。\n", n)
	} else {
		fmt.Printf("方程x^2 + y^2 = %d的所有解为: \n", n)
		for _, sol := range solutions {
			fmt.Printf("x = %d, y = %d\n", sol.x, sol.y)
		}
	}
}

func main() {
	var n int
	fmt.Print("请输入一个整数n: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	findSolutions(n)
}
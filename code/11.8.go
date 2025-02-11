package main

import (
	"fmt"
)

func extendedGCD2(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := extendedGCD2(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return g, x, y
}

func chineseRemainderTheorem(b, m, c, n int) int {
	g, x1, x2 := extendedGCD2(m, n)
	if g != 1 {
		fmt.Println("需要m和n互质。")
		return -1
	}

	x := (b*n*x2 + c*m*x1) % (m * n)
	if x < 0 {
		x += m * n
	}

	return x
}

func main() {
	var b, m, c, n int
	fmt.Print("请输入整数b, m, c, n: ")
	_, err := fmt.Scan(&b, &m, &c, &n)
	if err != nil {
		return
	}

	x := chineseRemainderTheorem(b, m, c, n)
	if x != -1 {
		fmt.Printf("满足条件的解为: x = %d\n", x)
	}
}
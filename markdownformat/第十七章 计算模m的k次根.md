>**如何计算模$m$的$k$次根**：设$b$，$k$与$m$是已知整数，满足$$gcd(b, m) = 1 \ 与 \ gcd(k, \sigma(m)) = 1$$下述步骤给出同余式$$x^k \equiv b \pmod{m}$$的解。
>	1. 计算$\varphi(m)$。
>	2. 求满足$ku - \varphi(m)v = 1$的正整数$u$与$v$。
>	3. 用逐次平方法计算$b^u \pmod{m}$。

>**习题**
- *17.1* 解同余式$x^{329} \equiv 452 \pmod{1147}$。
- *17.2*
	- *(a)* 解同余式$x^{113} \equiv 347 \pmod{463}$。
	- *(b)* 解同余式$x^{275} \equiv 139 \pmod{588}$。
- *17.3*
	- *(a)* 设$b$，$k$与$m$是整数，满足$gcd(b, m) = 1 \ 与 \ gcd(k, \varphi(m)) = 1$，证明$b$恰好有一个模$m$的$k$次根。
	- *(b)* 假设换成$gcd(k, \varphi(m)) > 1$。证明$b$没有模$m$的$k$次根，或者$b$至少有两个模$m$的$k$次根。
	- *(c)* 如果$m = p$是素数，观察例子，试求$b$模$p$的$k$次根的个数的公式。
- *17.4* 求解$x^k \equiv b \pmod{m}$的方法是：首先求满足$ku - \varphi(m)v = 1$的整数$u$与$v$，然后可知解为$x \equiv b^u \pmod{m}$。然而，我们仅证明如果$gcd(b, m) = 1$，则方法可行，因为可使用欧拉公式$b^{\varphi(m)} \equiv 1 \pmod{m}$。
	- *(a)* 如果$m$是不同素数的乘积，证明即使$gcd(b, m) > 1$，$x \equiv b^u \pmod{m}$的解。
	- *(b)* 证明我们的方法对同余式$x^5 \equiv 6 \pmod{9}$不可行。
- *17.5*
	- *(a)* 试用本章的方法计算$23$模$1279$的平方根。为什么会出错呢？
	- *(b)* 一般地，如果$p$是奇素数，解释为什么本章的方法不能用来求模$p$的平方根。
	- *(c)* 更一般地，如果$gcd(k, \varphi(m))$大于$1$，解释为什么计算模$m$的$k$次根方法不可行。
- *17.6* 编写求解$x^k \equiv b \pmod{m}$的程序。为用户提供用来计算$\varphi(m)$的因数分解$m$的选择。
```go
package main  
  
import (  
   "fmt"  
   "math"
)  
  
func factorization(n int) map[int]int {  
   factors := make(map[int]int)  
   for i := 2; i <= int(math.Sqrt(float64(n))); i++ {  
      for n%i == 0 {  
         factors[i]++  
         n /= i  
      }   
   }  
   if n > 1 {  
      factors[n] = 1  
   }  
   return factors  
}  
  
func euler(m int) int {  
   factors := factorization(m)  
   phi := m  
   for p := range factors {  
      phi -= phi / p  
   }  
   return phi  
}  
  
func extendedGCD(a, b int) (int, int, int) {  
   x0, x1, y0, y1 := 1, 0, 0, 1  
   for b != 0 {  
      q := a / b  
      a, b = b, a%b  
      x0, x1 = x1, x0-q*x1  
      y0, y1 = y1, y0-q*y1  
   }  
   return a, x0, y0  
}  
  
func modExp(base, exp, mod int) int {  
   result := 1  
   base = base % mod  
   for exp > 0 {  
      if exp%2 == 1 {  
         result = (result * base) % mod  
      }      
      exp = exp >> 1  
      base = (base * base) % mod  
   }  
   return result  
}  
  
func modKthRoot(k, b, m int) (int, error) {  
   phiM := euler(m)  
  
   _, u, _ := extendedGCD(k, phiM)  
   if u < 0 {  
      u += phiM  
   }  
  
   root := modExp(b, u, m)  
   return root, nil  
}  
  
func main() {  
   k := 329  
   b := 452  
   m := 1147  
  
   root, err := modKthRoot(k, b, m)  
   if err != nil {  
      fmt.Println("错误: ", err)  
   } else {  
      fmt.Printf("x^%d ≡ %d (mod %d) 的解为 x = %d\n", k, b, m, root)  
   }  
}
```
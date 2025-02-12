>与一个平方数模$p$同余的不是$p$倍数的数称为模$p$的二次剩余。
>不与任何一个平方数模$p$同余的数称为模$p$的二次非剩余。

>设$p$为一个奇素数，则恰有${p - 1} \over 2$个模$p$的二次剩余，且恰有${p - 1} \over 2$个模$p$的二次非剩余。

>**二次剩余乘法法则1**：设$p$为奇素数，则
>1. 两个模$p$的二次剩余的积是二次剩余。
>2. 二次剩余与二次非剩余的积是二次非剩余。
>3. 两个二次非剩余的积是二次剩余。

>$a$模$p$的勒让德符号是$$({a \over p}) = \left\{\begin{aligned} &1 \ \ \ \ \ 若a是模p的二次剩余 \\ -&1 \ \ \ \ \ 若a是模p的二次非剩余 \end{aligned} \right.$$

>**二次剩余乘法法则2**：设$p$为奇素数，则$$({a \over p})({b \over p}) = ({ab \over p})$$

>**习题**
- *20.1* 列出模$19$的所有二次剩余与二次非剩余。
- *20.2* 对每个奇素数$p$，考虑下面两个数$$\begin{aligned} A &= 模p的所有满足1 \leq a < p的二次剩余a的和 \\ B &= 模p的所有满足1 \leq a < p的二次非剩余a的和 \end{aligned}$$
	- *(a)* 对所有满足$p < 20$的奇素数列出$A$和$B$。
	- *(b)* $A+B$的值等于多少？证明你的猜测。
	- *(c)* 计算$A \pmod{p}$和$B \pmod{p}$。寻找模式并证明之。
	- *(d)* 收集更多的资料并给出判别$p$是否满足$A = B$的准则。
	- *(e)* 编写一个计算机程序计算$A$与$B$，并用它对奇素数$p \leq 100$制作表格。若$A \neq B$，$A$与$B$哪一个会更大呢？是$A$还是$B$？
```go
package main  
  
import (  
   "fmt"  
   "math"
)  
  
func isQuadraticResidue(a, p int) bool {  
   for x := 1; x < p; x++ {  
      if (x * x) % p == a {  
         return true  
      }   
   }  
   return false  
}  
  
func calculateAandB(p int) (int, int) {  
   var A, B int  
   for a := 1; a < p; a++ {  
      if isQuadraticResidue(a, p) {  
         A += a  
      } else {  
         B += a  
      }   
   }  
   return A, B  
}  
  
func isPrime(n int) bool {  
   if n <= 1 {  
      return false  
   }  
   for i := 2; i <= int(math.Sqrt(float64(n))); i++ {  
      if n%i == 0 {  
         return false  
      }   
   }  
   return true  
}  
  
func main() {  
   var primes []int  
   for p := 3; p <= 100; p++ {  
      if isPrime(p) {  
         primes = append(primes, p)  
      }  
   }  
  
   fmt.Printf("p: ")  
   for _, p := range primes {  
      fmt.Printf("%-5d", p)  
   }  
   fmt.Println()  
  
   fmt.Printf("A: ")  
   for _, p := range primes {  
      A, _ := calculateAandB(p)  
      fmt.Printf("%-5d", A)  
   }  
   fmt.Println()  
   fmt.Printf("B: ")  
   for _, p := range primes {  
      _, B := calculateAandB(p)  
      fmt.Printf("%-5d", B)  
   }  
   fmt.Println()  
}
```
![[AandB.png]]
- *20.3* 数$a$称为模$p$的三次剩余，是指它与一个立方数模$p$同余，也就是说，存在一个数$b$使得$a \equiv b^3 \pmod{p}$。
	- *(a)* 对$p = 5，7，11，13$，列出模$p$的所有三次剩余。
	- *(b)* 找出两个数$a_1$和$b_1$，使$a_1，b_1$都不是模$19$的三次剩余，但$a_1b_1$是模$19$的三次剩余。类似地，找出两个数$a_2$和$b_2$，使$a_2，b_2，a_2b_2$都不是模$19$的三次剩余。
	- *(c)* 若$p \equiv 2 \pmod{3}$，对哪些$a$是三次剩余做一个猜测并证明之。
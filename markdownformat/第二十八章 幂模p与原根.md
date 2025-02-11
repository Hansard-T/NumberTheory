>**次数整除性质**：设$a$是不被素数$p$整除的整数，假设$a^n \equiv 1 \pmod{p}$，则次数$e_p(a)$整除$n$。特别地，次数$e_p(a)$总整除$p-1$。

>**原根定理**：每个素数$p$都有原根。更精确地，有恰好$\phi(p - 1)$个模$p$的原根。

>**习题**
- *28.1* 设$p$是素数。
	- *(a)* $1 + 2 + 3 + ··· + (p - 1) \pmod{p}$的值是什么？
	- *(b)* $1^2 + 2^2 + 3^2 + ··· + (p - 1)^2 \pmod{p}$的值是什么？
	- *(c)* 对任何正整数$k$，求$1^k + 2^k + 3^k + ··· + (p - 1)^k \pmod{p}$的值，证明你的答案是正确的。
- *28.2* 对任何整数$a$与$m$，$gcd(a, m) = 1$，设$e_m(a)$是使得$a^e \equiv 1 \pmod{m}$的最小指数$(e \geq 1)$。我们称$e_m(a)$为$a$模$m$的次数。
	- *(a)* 计算下述$e_m(a)$的值：*(i)* $e_9(2)$ *(ii)* $e_{15}(2)$ *(iii)* $e_{16}(3)$ *(iv)* $e_{10}(3)$
	- *(b)* 证明$e_m(a)$总是整除$\phi(m)$。
- *28.3* 在本习题中将研究奇数$m$的$e_m(2)$的值。为节省空间，将$e_m(2)$写成$e_m$，因此对于本习题，$e_m$是模$m$余$1$的$2$的最小幂次。
	- *(a)* 对每个奇数$11 \leq m \leq 19$，计算$e_m$的值。
	- *(b)*  猜测用$e_m$与$e_n$表示的$e_{mn}$的公式。
	- *(c)* 用*(b)*中猜测的公式求$e_{11227}$的值。(注意$11227 = 103 \cdot 109$)
	- *(d)* 证明你在*(b)*中猜测的公式是正确的。
	- *(e)* 由表格猜测用$e_p$，$p$，$k$表示的$e_{p^k}$的公式是正确的吗？
	- *(f)* 你能证明*(e)*中猜测的$e_{p^k}$的公式是正确的吗？
- *28.4*
	- *(a)* 求模$13$的所有原根。
	- *(b)* 对整除$12$的每个整数$d$，列出满足$1 \leq a < 13$与$e_{13}(a) = d$的$a$。
- *28.5*
	- *(a)* 如果$g$是模$37$的原根，则数$g^2，g^3，···，g^8$中哪些是模$37$的原根？
	- *(b)* 如果$g$是模$p$的原根，给出一种容易使用的法则，以确定$g^k$是否是模$p$的原型，并证明你的法则是正确的。
	- *(c)* 假设$g$是模素数$p = 21169$的原根。使用*(b)*中的法则来确定$g^2，g^3，···，g^{20}$中的哪些数是模$21169$的原根。
- *28.6*
	- *(a)* 求小于$20$的所有素数，使得$3$是其原根。
	- *(b)* 如果你了解编写计算机程序的原理，求小于$100$的所有素数，使得$3$是其原根。
```go
package main  
  
import "fmt"  
  
func modPow(a, e, p int) int {  
   result := 1  
   a = a % p  
  
   for e > 0 {  
      if e % 2 == 1 {  
         result = (result * a) % p  
      }      
      a = (a * a) % p  
      e = e / 2  
   }  
  
   return result  
}  
  
func isPrimitiveRoot(p int) bool {  
   for e := 1; e < p; e++ {  
      if modPow(3, e, p) == 1 {  
         if e == p - 1 {  
            return true  
         }  
         return false  
      }   
   }  
  
   return false  
}  
  
func main() {  
   primes20 := []int{2, 3, 5, 7, 11, 13, 17, 19}  
   primes100 := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}  
   var result20 []int  
   var result100 []int  
  
   for _, p := range primes20 {  
      if p == 3 {  
         continue  
      }  
      if isPrimitiveRoot(p) {  
         result20 = append(result20, p)  
      }  
   }  
  
   for _, p := range primes100 {  
      if p == 3 {  
         continue  
      }  
      if isPrimitiveRoot(p) {  
         result100 = append(result100, p)  
      }  
   }  
  
   fmt.Printf("3是原根的小于20的素数有: %v\n", result20)  
   fmt.Printf("3是原根的小于100的素数有: %v\n", result100)  
}
```
![[28.6.png]]
- *28.7* 如果$a = p^2$是完全平方数，$p$是奇素数，解释$a$不可能是模$p$的原根的原因。
- *28.8* 设$p$是奇素数，$g$是模$p$的原根。
	- *(a)* 证明$g^k$是模$p$的二次剩余当且仅当$k$是偶数。
	- *(b)* 利用*(a)*给出下述结论的一个快速证明：两个二次非剩余的积是二次剩余，更一般地，$({a \over p})({b \over p}) = ({ab \over p})$。
	- *(c)* 利用*(a)*给出欧拉准则$a^{(p - 1)/2} \equiv ({a \over p}) \pmod{p}$的一个快速证明。
- *28.9* 假设$q$是模$4$余$1$的素数，且$p = 2q + 1$也是素数。证明$2$是模$p$的原根。
- *28.10* 设$p$是一个素数，$k$是一个不被$p$整除的数，$b$是一个在模$p$下有$k$次方根的数。求出$b$模$p$下的$k$次方根的个数公式，并证明该公式是正确的。(提示：你的公式应该只依赖于$p$和$k$，而不依赖于$b$)
- *28.11* 编写程序计算$e_p(a)$，这是使得$a^e \equiv 1 \pmod{p}$的最小正指数$e$。
```go
package main  
  
import (  
   "fmt"  
   "math"
)  
  
func ep(a, p int) int {  
   for e := 1; e < p; e++ {  
      ae := math.Pow(float64(a), float64(e))  
      if int(ae) % p == 1 {  
         return e  
      }   
   }  
   return -1  
}  
  
func main() {
   a := 7  
   p := 11  
   fmt.Printf("e_%d(%d) = %d\n", p, a, ep(a, p))  
}
```
- *28.12* 编写程序，求给定素数$p$的最小原根。制作$100$与$200$之间以$2$为原根的所有素数的列表。
```go
package main  
  
import (  
   "fmt"  
   "math"
)  
  
func minprimeroot(p int) int {  
   for i := 2; i < p; i++ {  
      ipsub1 := int(math.Pow(float64(i), float64(p-1))) % p  
      if ipsub1 == 1 {  
         return i  
      }   }  
   return -1  
}  
  
func modPow(a, e, p int) int {  
   result := 1  
   a = a % p  
  
   for e > 0 {  
      if e % 2 == 1 {  
         result = (result * a) % p  
      }      
      a = (a * a) % p  
      e = e / 2  
   }  
  
   return result  
}  
  
func isPrimitiveRoot(p int) bool {  
   for e := 1; e < p; e++ {  
      if modPow(2, e, p) == 1 {  
         if e == p - 1 {  
            return true  
         }  
         return false  
      }   
   }  
  
   return false  
}  
  
func main() {  
   var p int  
   fmt.Printf("请输入一个素数: ")  
   _, err := fmt.Scan(&p)  
   if err != nil {  
      return  
   }  
  
   x := minprimeroot(p)  
   fmt.Printf("素数%v的最小原根: %v\n", p, x)  
  
   primes200 := []int{101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}  
   var result200 []int  
   for _, p := range primes200 {  
      if isPrimitiveRoot(p) {  
         result200 = append(result200, p)  
      }  
   }  
  
   fmt.Printf("100与200之间以2为原根的所有素数: %v", result200)  
}
```
![[28.12.png]]
- *28.13* 如果$a$与$m$，$n$都互素，且$gcd(m, n) = 1$，求用$e_m(a)$与$e_n(a)$表示的$e_{mn}(a)$的公式。
- *28.14* 对任何数$m \geq 2$(不必是素数)，如果同余于$1 \pmod{m}$的$g$的最小幂次是$\phi(m)$，我们称$g$是模$m$的原根。换句话说，如果$gcd(g, m) = 1$且对所有幂$1 \leq k < \phi(m)$有$g^k \equiv 1 \pmod{m}$，则$g$是模$m$的原根。
	- *(a)* 对每个数$2 \leq m \leq 25$，确定是否存在模$m$的原根。
	- *(b)* 使用*(a)*的数据，猜测哪些$m$有原根，哪些$m$没有原根。
	- *(c)* 证明*(b)*中的猜测是正确的。
- *28.15* 回顾置换阵的定义，它是一个每行、每列都恰有一个点的阵列。
	- *(a)* $N \times N$置换阵有多少个？
	- *(b)* 我们可以通过将点用$1$替代而其余位置用$0$替代的方法将一个点阵变成一个矩阵。$A^6$的值是什么？
	- *(c)* 设$A$是一个$N \times N$置换矩阵，既是由置换矩阵得到的矩阵。证明：存在整数$k \geq 1$，使得$A^k$为单位矩阵。
	- *(d)* 找出满足下述条件的$N \times N$置换矩阵$A$的一个例子：使$A^k$为单位矩阵的最小正整数$k$满足$k > N$。
- *28.16*
	- *(a)* 找出所有$3$阶$Costas$阵列。
	- *(b)* 写出一个$4$阶$Costas$阵列。
	- *(c)* 写出一个$5$阶$Costas$阵列。
	- *(d)* 写出一个$7$阶$Costas$阵列。
- *28.17* 利用$Welch$构造找出一个$16$阶$Costas$阵列。当然，要表明你使用了哪一个原根。
- *28.18* 对于产生$p - 2$阶$Costas$阵列的$Lempel$与$Golumb$构造，本习题描述了一个特殊情形。
	- *(a)* 设$g_1$与$g_2$是模$p$的原根。证明：对每个$1 \leq i \leq p - 2$，存在唯一的$1 \leq j \leq p - 2$满足$$g_1^i + g_2^j = 1$$
	- *(b)* 构造一个$(p - 2) \times (p - 2)$阵列如下：如果$i$，$j$满足$g_1^i + g_2^j = 1$，则将一个点放入第$i$行、第$j$列。证明所得阵列是$Costas$阵列。
	- *(c)* 利用$Lempel - Golumb$构造写出两个$15$阶$Costas$阵列。对第一个阵列，使用$g_1 = g_2 = 5$，对第二个阵列，使用$g_1 = 3$，$g_2 = 6$。
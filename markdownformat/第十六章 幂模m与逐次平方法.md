>**逐次平方法计算$a^k \ (mod \ m)$**：用下述步骤计算$a^k \ (mod \ m)$的值：
>1. 将$k$表成$2$的幂次和：$$k = u_0 + u_1 \cdot 2 + u_2 \cdot 2^2 + u_3 \cdot 2^3 + ··· + u_r \cdot 2^r$$其中每个$u_i$是$0$或$1$。
>2. 使用逐次平方法制作模$m$的$a$的幂次表。$$\begin{aligned} a^1  &\equiv A_0 \pmod{m} \\ a^2  &\equiv (a^1)^2 \equiv A_0^2 \equiv A_1 \pmod{m} \\ a^4  &\equiv (a^2)^2 \equiv A_1^2 \equiv A_2 \pmod{m} \\ a^8  &\equiv (a^4)^2 \equiv A_2^2 \equiv A_3 \pmod{m} \\ &\vdots \\ a^{2^r} &\equiv (a^{2^{r-1}})^2 \equiv A_{r-1}^2 \equiv A_r \pmod{m} \end{aligned}$$注意要计算表的每一行，仅需要取前一行最末的数，平方它然后用模$m$简化。也注意到表有$r + 1$行，其中$r$是第$1$步中$k$的二进制展开式中$2$的最高指数。
>3. 乘积$$A_0^{u_0} \cdot A_1^{u_1} \cdot A_2^{u_2} ··· A_r^{u_r} \pmod{m}$$同余于$a^k \pmod{m}$。注意到所有$u_i$是$0$或$1$，因此这个数实际上是$u_i = 1$的那些$A_i$的乘积。

>**习题**
- *16.1* 使用逐次平方法计算下述幂。
	- *(a)* $5^{13} \pmod{23}$
	- *(b)* $28^{749} \pmod{1147}$
- *16.2* 本章讲述的逐次平方法可以相当有效地计算$a^k \pmod{m}$，但它确实包含创建模$m$的$a$的幂次的表格。
	- *(a)* 使用你选取的计算机语言，在计算机上完成更有效的进行逐次平方的算法程序。
	- *(b)* 使用你的程序计算下述值：*(i)* $2^{1000} \pmod{2379}$ *(ii)* $567^{1234} \pmod{4321}$ *(iii)* $47^{258008} \pmod{1315171}$
```go
package main  
  
import "fmt"  
  
func sucsquare(a, k, m int) int {  
   b := 1  
  
   for k >= 1 {  
      if k % 2 == 1 {  
         b = (a * b) % m  
      }  
      a = (a * a) % m  
      k = k / 2  
   }  
  
   return b  
}  
  
func main() {  
   var a, k, m int  
   fmt.Print("请输入整数a, k, m: ")  
   _, err := fmt.Scan(&a, &k, &m)  
   if err != nil {  
      return  
   }  
  
   result := sucsquare(a, k, m)  
   fmt.Printf("%d^%d mod %d = %d\n", a, k, m, result)  
}
```
- *16.3*
	- *(a)* 用逐次平方法计算$7^{7386} \pmod{7387}$。$7387$是素数吗？
	- *(b)* 用逐次平方法计算$7^{7392} \pmod{7393}$。$7393$是素数吗？
- *16.4* 编写程序验证$n$是合数还是可能素数。
```go
package main  
  
import (  
   "fmt"  
   "math/rand"   
   "time"
)  
  
func sucsquare(a, b, m int) int {  
   result := 1  
  
   for b > 0 {  
      if b%2 == 1 {  
         result = (result * a) % m  
      }      
      a = (a * a) % m  
      b = b / 2  
   }  
  
   return result  
}  
  
func checkPrime(n int) {  
   if n <= 1 {  
      fmt.Printf("%d是合数", n)  
      return   
   }  
  
   if n == 2 {  
      fmt.Printf("%d是质数", n)  
      return
   }  
  
   rand.Seed(time.Now().UnixNano())  
  
   checks := 10  
   for i := 0; i < checks; i++ {  
      a := rand.Intn(n-2) + 2  
  
      result := sucsquare(a, n-1, n)  
  
      if result != 1 {  
         fmt.Printf("%d是合数", n)  
         return
	  }  
   }  
  
   fmt.Printf("%d可能是素数", n)  
}  
  
func main() {  
   var n int  
   fmt.Print("请输入一个整数: ")  
   _, err := fmt.Scan(&n)  
   if err != nil {  
      return  
   }  
  
   checkPrime(n)  
}
```
- *16.5* 用逐次平方法计算$2^{9990} \pmod{9991}$，再用你的答案说明你是否相信$9991$是素数。
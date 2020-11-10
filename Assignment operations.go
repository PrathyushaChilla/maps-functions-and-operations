program1

package main
import "fmt"
func main() {
    var a int =567
    var b int =341

    fmt.Printf("Assignment Operators")
    a+=b
    fmt.Printf("a+=b :%d\n", a)
    a-=b
    fmt.Printf("a-=b :%d\n", a)
    a*=b
    fmt.Printf("a*=b :%d\n", a)
    a/=b
    fmt.Printf("a/=b :%d\n", a)
    a%=b
    fmt.Printf("a%%=b :%d\n", a) 
}
output:
Assignment Operatorsa+=b :908
a-=b :567
a*=b :193347
a/=b :567
a%=b :226




program2

package main
import "fmt"
func main() {
var x, y = 31, 56
x = y
fmt.Println("= ", x)

x = 12
x += y
fmt.Println("+=", x)
x = 55
x -= y
fmt.Println("-=", x)
x = 4
x *= y
fmt.Println("*=", x)
x = 100
x /= y
fmt.Println("/=", x)
x = 60
x %= y
fmt.Println("%=", x)
}
output:
=  56
+= 68
-= -1
*= 224
/= 1
%= 4

Program exited





program3

package main
import "fmt"
func main() {
    var a int =30
    var b int =5

    fmt.Printf("Go Assignment Operators")
    a+=b
    fmt.Printf("a+=b :%d\n", a)
    a-=b
    fmt.Printf("a-=b :%d\n", a)
    a*=b
    fmt.Printf("a*=b :%d\n", a)
    a/=b
    fmt.Printf("a/=b :%d\n", a)
    a%=b
    fmt.Printf("a%%=b :%d\n", a)
     }
Output:
Go Assignment Operators
a+=b :35
a-=b :30
a*=b :150
a/=b :30
a%=b :0




program4

package main
import "fmt"
func main() {
   var a int = 21
   var c int
   c =  a
   fmt.Printf("Line 1 - =  Operator Example, Value of c = %d\n", c )
   c +=  a
   fmt.Printf("Line 2 - += Operator Example, Value of c = %d\n", c )
   c -=  a
   fmt.Printf("Line 3 - -= Operator Example, Value of c = %d\n", c )
   c *=  a
   fmt.Printf("Line 4 - *= Operator Example, Value of c = %d\n", c )
   c /=  a
   fmt.Printf("Line 5 - /= Operator Example, Value of c = %d\n", c )
   c  = 200; 
   c <<=  2
   fmt.Printf("Line 6 - <<= Operator Example, Value of c = %d\n", c )
   c >>=  2
   fmt.Printf("Line 7 - >>= Operator Example, Value of c = %d\n", c )
   c &=  2
   fmt.Printf("Line 8 - &= Operator Example, Value of c = %d\n", c )
   c ^=  2
   fmt.Printf("Line 9 - ^= Operator Example, Value of c = %d\n", c )
   c |=  2
   fmt.Printf("Line 10 - |= Operator Example, Value of c = %d\n", c )
}
Output:
Line 1 - =  Operator Example, Value of c = 21
Line 2 - += Operator Example, Value of c = 42
Line 3 - -= Operator Example, Value of c = 21
Line 4 - *= Operator Example, Value of c = 441
Line 5 - /= Operator Example, Value of c = 21
Line 6 - <<= Operator Example, Value of c = 800
Line 7 - >>= Operator Example, Value of c = 200
Line 8 - &= Operator Example, Value of c = 0
Line 9 - ^= Operator Example, Value of c = 2
Line 10 - |= Operator Example, Value of c = 2





program5

package main 
import "fmt"
func main() { 
var p int = 45
var q int = 50
p = q 
fmt.Println(p) 
p += q 
fmt.Println(p) 
p -= q
fmt.Println(p) 
p*= q 
fmt.Println(p) 
p /= q 
fmt.Println(p) 
 p %= q 
 fmt.Println(p)     
}
 
Output:
50
100
50
2500
50
0
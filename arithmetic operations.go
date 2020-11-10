program1

package main
import "fmt"
func main() {
var a int=10
var b int=100
fmt.Println("Arithmetic Operators")
fmt.Println(a+b)
fmt.Println(a-b)
fmt.Println(a*b)
fmt.Println(a/b)
fmt.Println(a%b)
}
output:
Arithmetic Operators
110
-90
1000
0
10




program2

package main
import "fmt"
func main() {
   var a int = 29
   var b int = 30
   var c int
   c = a + b
   fmt.Printf("Line 1 - Value of c is %d\n", c )
   c = a - b
   fmt.Printf("Line 2 - Value of c is %d\n", c )
   c = a * b
   fmt.Printf("Line 3 - Value of c is %d\n", c )
   c = a / b
   fmt.Printf("Line 4 - Value of c is %d\n", c )
    c = a % b
   fmt.Printf("Line 5 - Value of c is %d\n", c )
   a++
   fmt.Printf("Line 6 - Value of a is %d\n", a )
   a--
   fmt.Printf("Line 7 - Value of a is %d\n", a )
}
output:
Line 1 - Value of c is 59
Line 2 - Value of c is -1
Line 3 - Value of c is 870
Line 4 - Value of c is 0
Line 5 - Value of c is 29
Line 6 - Value of a is 30
Line 7 - Value of a is 29
Program exited.





program3

package main 
import "fmt"
func main() { 
p:= 3
q:= 6
result1:= p + q 
fmt.Printf("Result of p + q = %d", result1) 
result2:= p - q 
fmt.Printf("\nResult of p - q = %d", result2) 
result3:= p * q 
fmt.Printf("\nResult of p * q = %d", result3) 
result4:= p / q 
fmt.Printf("\nResult of p / q = %d", result4) 
result5:= p % q 
fmt.Printf("\nResult of p %% q = %d", result5) 
} 
output:
Result of p + q = 9
Result of p - q = -3
Result of p * q = 18
Result of p / q = 0
Result of p % q = 3
Program exited.
program1

package main 
import "fmt"
func main() { 
p:= 29
q:= 30
result1:= p == q 
fmt.Println(result1) 
result2:= p != q 
fmt.Println(result2) 
result3:= p < q 
fmt.Println(result3) 
result4:= p > q 
fmt.Println(result4) 
result5:= p >= q 
fmt.Println(result5) 
result6:= p <= q 
fmt.Println(result6)     
} 
output:
false
true
true
false
false
true



program2

import "fmt"
func main() {
var a int = 21
var b int = 10
if( a == b ) 
{
fmt.Printf("Line 1 - a is equal to b\n" )
   } 
else 
{
fmt.Printf("Line 1 - a is not equal to b\n" )
 }
 if ( a < b )
 {
 fmt.Printf("Line 2 - a is less than b\n" )
 } 
else 
{
 fmt.Printf("Line 2 - a is not less than b\n" )
   } 
   if ( a > b ) {
      fmt.Printf("Line 3 - a is greater than b\n" )
   } else {
      fmt.Printf("Line 3 - a is not greater than b\n" )
   }
   a = 5
   b = 20
   if ( a <= b ) {
      fmt.Printf("Line 4 - a is either less than or equal to  b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("Line 5 - b is either greater than  or equal to b\n" )
   }
}
output:
Line 1 - a is not equal to b
Line 2 - a is not less than b
Line 3 - a is greater than b
Line 4 - a is either less than or equal to  b
Line 5 - b is either greater than  or equal to b

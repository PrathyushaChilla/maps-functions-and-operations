program1
package main
import "fmt"
func main() { 
   var a int = 29
   var b int = 30
   var ret int
   ret = max(a, b)
   fmt.Printf( "Max value is : %d\n", ret )
}
func max(num1, num2 int) int{
   var result int
   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}
output:
Max value is : 30
Program exited.



program2

package main
import "fmt"
func swap(x, y string) (string, string) {
   return y, x
}
func main() {
   a, b := swap("prathyusha", "reddy")
   fmt.Println(a, b)
}
output:
reddy prathyusha
Program exited.





program3

package main
import "fmt" 
func area(length, width int)int{
     
    Ar := length* width
    return Ar
}
func main() {
   
   fmt.Printf("Area of rectangle is : %d", area(29, 80))
}
output:
Area of rectangle is : 2320
Program exited.




program4
package main
import (  
    "fmt"
)

func appendStr() func(string) string {  
    t := "Good"
    c := func(b string) string {
        t = t + " " + b
        return t
    }
    return c
}

func main() {  
    a := appendStr()
    b := appendStr()
    fmt.Println(a("Morning"))
    fmt.Println(b("Everyone"))

    fmt.Println(a("prathyu"))
    fmt.Println(b("!"))
}
output:
Good Morning
Good Everyone
Good Morning prathyu
Good Everyone !
Program exited.



program5
package main

import "fmt"

func inc(x int) int {
    x++
    return x
}

func dec(x int) int {
    x--
    return x
}

func apply(x int, f func(int) int) int {

    r := f(x)
    return r
}

func main() {
    r1 := apply(12, inc)
    r2 := apply(4, dec)
    fmt.Println(r1)
    fmt.Println(r2)
}
output
13
3


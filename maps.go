program1

package main
import (
	"fmt"
)
func main() {
statepopulation := map[string]int{
"andhra pradesh": 3234,
"tamil nadu": 56789,
"kerala": 12345,
}
	fmt.Println(statepopulation)
}

output:
map[andhra pradesh:3234 kerala:12345 tamil nadu:56789]




program2

package main
import (
	"fmt"
)
func main() {
statepopulation := make(map[string]int)
statepopulation = map[string]int{
"andhra pradesh": 3234,
"tamil nadu": 56789,
"kerala": 12345,
}
	fmt.Println(statepopulation["kerala"])
}
output:
12345
Program exited.




program3

package main
import (
	"fmt"
)
func main() {
statepopulation := make(map[string]int)
statepopulation = map[string]int{
"andhra pradesh": 3234,
"tamil nadu": 56789,
"kerala": 12345,
}
statepopulation["karnataka"] = 45632
	fmt.Println(statepopulation["karnataka"])
}
output:
45632
Program exited.



program4

package main
import (
	"fmt"
)
func main() {
statepopulation := make(map[string]int)
statepopulation = map[string]int{
"andhra pradesh": 3234,
"tamil nadu": 56789,
"kerala": 12345,
}
statepopulation["karnataka"] = 45632
	fmt.Println(statepopulation)
}
output:
map[andhra pradesh:3234 karnataka:45632 kerala:12345 tamil nadu:56789]






program5

package main
import (
	"fmt"
)
func main() {
statepopulation := make(map[string]int)
statepopulation = map[string]int{
"andhra pradesh": 3234,
"tamil nadu": 56789,
"kerala": 12345,
}
sp := statepopulation
delete (sp,"kerala")
fmt.Println(sp)
	fmt.Println(statepopulation)
}
output:
map[andhra pradesh:3234 tamil nadu:56789]
map[andhra pradesh:3234 tamil nadu:56789]





program6

package main
import (
	"fmt"
)
type Doctor struct {
number int
actorName string
companions []string
}
func main() {
aDoctor := Doctor{
number : 3,
actorName: "prathyusha",
companions: []string {
"supraja",
"damu",
},
}
fmt.Println(aDoctor.companions)
}
output:
[supraja damu]

Program exited.





program7

package main
import (
	"fmt"
)
func main() {
aDoctor := struct{Name string} {Name: "Prathyusha"}
fmt.Println(aDoctor)
}
output:
{Prathyusha}

Program exited.
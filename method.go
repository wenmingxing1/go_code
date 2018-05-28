/*方法*/

package main 
import "fmt"

type TwoInts struct {
	a int 
	b int
}

func main() {
	two1 := TwoInts{12, 10}

	fmt.Printf("The sum is : %d\n", two1.AddThem())
	fmt.Printf("Add them to the param : %d\n", two1.AddToParam(20))

}

func (this *TwoInts) AddThem() int {
	return this.a + this.b
}

//含参method
func (this *TwoInts) AddToParam(x int) int {
	return this.AddThem() + x
	//return this.a + this.b + x
}
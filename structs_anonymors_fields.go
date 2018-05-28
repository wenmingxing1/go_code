/*匿名字段*/

package main 
import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32	
	int    //anoymous field
	innerS   //anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5		//很像继承
	outer.in2 = 10

	fmt.Println(outer)
}
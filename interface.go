/*接口，即go中实现多态的方式*/

package main 
import (
	"fmt"
	"math"
)

//定义接口struct
type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	height, width float64
}

func (r *rect) area() float64 {
	return r.height*r.width
}

func (r *rect) perimeter() float64 {
	return (r.width+r.height)*2
}


type circle struct {
	redius float64
}

func (c *circle) area() float64 {
	return math.Pi*c.redius*c.redius
}

func (c *circle) perimeter() float64 {
	return 2*math.Pi*c.redius
}

func main() {
	r := rect{19,20}
	c := circle{5}

	s := []shape{&r,&c}	//通过指针实现

	for _,sh:=range s {
		fmt.Println(sh)
		fmt.Println(sh.area())
		fmt.Println(sh.perimeter())
	}

}
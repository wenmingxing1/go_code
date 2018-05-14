/*结构体方法*/

package main 
import "fmt"

type rect struct {
	height,width int
}

//定义结构体方法
func (r *rect) area() int {
	return r.height*r.width
}

func (r *rect) perimeter() int {
	return 2*(r.height+r.width)
}

func main() {
	r := rect{10,15}

	fmt.Println(r.area())

	fmt.Println(r.perimeter())

	pr := &r;
	pr.height = 20
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
}


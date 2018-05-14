/*函数不定参*/

package main 
import "fmt"

func main() {
	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3))

	a := []int{1,2,3,4,5}
	fmt.Println(sum(a...))	//传入数组

}

//不定参函数的定义
func sum(nums ...int) int {
	total := 0
	//以_占位
	for _,num:= range nums {
		total += num
	}
	return total
}
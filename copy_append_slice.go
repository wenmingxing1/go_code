/*slice的复制与追加*/

package main 
import "fmt"

func main() {
	s1_from := []int{1,2,3}
	s1_to := make([]int, 10)

	n := copy(s1_to, s1_from)	//返回copy成功元素的个数

	fmt.Println(s1_to)	//[1 2 3 0 0 0 0 0 0 0]
	fmt.Println(s1_from)
	fmt.Println(n)

	s13 := []int{1,2,3}
	s13 = append(s13,4,5,6)
	fmt.Println(s13)
}
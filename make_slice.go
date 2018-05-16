/*使用make创建slice*/

package main 
import "fmt"

func main() {
	//在使用make创建一个slice的过程中，其实也创建了相关数组，然后用slice指向了数组
	var slice1 []int = make([]int, 10)

	for i:= 0; i<len(slice1); i++ {
		slice1[i] = 5*i
	}

	for i:=0; i<len(slice1); i++ {
		fmt.Printf("%d,", slice1[i])
	}

	fmt.Printf("\n")

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
/*数组切片*/

package main 
import "fmt"

func main() {
	a := [5]int{1,2,3,4,5}

	b:=a[2:4]	//a[2]与a[3],不包括a[4]
	fmt.Println(b)

	b = a[:4]	//a[0],a[1],a[2],a[3]
	fmt.Println(b)

	b = a[2:]	//a[2],a[3],a[4]
	fmt.Println(b)
}
/*切片重组*/

package main 
import "fmt"

func main() {
	s := make([]int, 5, 10)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	for i:=0; i<cap(s); i++ {
		s = s[0:i+1]	//reslice一位
		s[i] = i
	} 

	for _,val := range s {
		fmt.Println(val)
	}
}
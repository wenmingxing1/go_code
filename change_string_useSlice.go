/*string是不能被改变的，所以可以通过slice对其进行改变*/

package main 
import "fmt"

func main() {
	var s string = "Hello"

	//s[0] = 'C'	//这样会报错
	slice := []byte(s)

	slice[0] = 'C'	//通过切片修改
	s = string(slice)

	fmt.Println(s)
}
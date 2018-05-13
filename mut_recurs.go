/*两个函数相互调用，形成调用闭环*/
package main 
import "fmt"

func main() {
	fmt.Printf("%d is even? %t", 16,even(16))
	fmt.Println()
	fmt.Printf("%d is odd? %t", 17,odd(17))
}

func even(i int) bool {
	if i==0 {
		return true
	} else {
		return odd(RevSign(i)-1)
	}
}

func odd(i int) bool {
	if i==0 {
		return false
	} else {
		return even(RevSign(i)-1)
	}
}

func RevSign(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
package main 
import "fmt"

func main() {
	nextNumFunc := nextNum()  

	for i:=0; i<10; i++ {
		fmt.Println(nextNumFunc())
	}
}

//返回一个函数，从而形成闭包
//因为在nextNum函数中，存在i，j两个状态,并保持了这两个元素的状态
//这个函数返回Fibonacci数列的下一个值
func nextNum() func() int {
	i,j:=1,1
	return func() int {
		var tmp = i+j
		i,j = j,tmp
		return tmp
	}
}
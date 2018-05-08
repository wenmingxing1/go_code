package main 
import "fmt"

func main() {
	var n int16 = 34
	var m int32

	m = int32(n)	//强制转换类型，如果不进行强制转换，
	//则会报错，因为go语言是强类型语言

	fmt.Printf("32 bit int is:%d\n", m)
	fmt.Printf("16 bit int is:%d\n", n)
}
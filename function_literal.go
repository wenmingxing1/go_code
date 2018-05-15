package main 
import "fmt"

func main() {
	//匿名函数的两种调用方式
	//1、赋值于变量完成调用
	f := func(i,j int) int {return i+j}
	fmt.Println(f(1,2))

	//2、直接完成调用
	fmt.Println(func(i,j int) int {return i+j} (1,2))
	
}

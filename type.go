package main 

import "fmt"

type MyInt int 		//自己定义类型别名

func main() {
	var a MyInt = 1		//使用类型别名
	fmt.Println(a)
}
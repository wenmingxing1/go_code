package main 

import (
	"fmt"
	"./trans"	//导入具有init函数的包
)

var twoPi = 2 * trans.Pi   	//这里可以看出大小写的关系

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi)
}
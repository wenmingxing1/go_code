package main 

import (
	"fmt"
	"runtime"
	"os"
)

func main() {
	var goos string = runtime.GOOS    //在变量名后标记类型string
	fmt.Printf("The operating system is : %s\n", goos)
	path := os.Getenv("PATH")	//:=表示声明并赋值，相当于在变量名前面加了var声明
	fmt.Printf("Path is %s\n", path)
}
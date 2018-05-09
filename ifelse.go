package main 

import "fmt"

func main() {
	var first int = 10
	var cond int

	if first < 0 {
		fmt.Println("The number is less than zero")
	} else if first > 0 && first < 5 {
		fmt.Println("The number is between 0 and 5")
	} else {
		fmt.Println("The number is greater than 5")
	}

	//if语句中可以对变量进行初始化或赋值，作用域为整个ifelse生存空间
	if cond = 5; cond > 10 {
		fmt.Println("The number is greater than 10")
	} else {
		fmt.Println("The number is less than 10")
	}
}
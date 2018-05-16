package main 

import "fmt"

func main() {
	a := [5]int{1,3,5,7,9}

	var s []int = a[0:3]

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = s[1:]	//向后移动一位
	fmt.Println(s)
}
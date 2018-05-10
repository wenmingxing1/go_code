package main 

import "fmt"

func main() {
	x:=min(1,3,2,0)
	fmt.Printf("The min number is %d\n", x)
	slice := []int{7,9,3,5,1}		//通过slice传递变长参数
	x = min(slice...)  

	fmt.Printf("The min number in the slice is %d\n",x)
}

func min(s ...int) int {
	if len(s) == 0 {return 0}
	min := s[0]
	for _,v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
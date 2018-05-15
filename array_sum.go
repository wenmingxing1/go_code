package main 
import "fmt"

func main() {
	arr := [5]int{1,2,3,4,5}
	fmt.Println(Sum(&arr)) 

	fmt.Println(arr)
}

func Sum(a *[5]int) int {
	var sum int
	for _,i := range(a) {
		sum += i
	}

	a[0] =100
	return sum
}
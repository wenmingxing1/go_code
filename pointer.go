package main 

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("%d, its location in memory:%p\n", i1,&i1)	//%p为指针
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
/*将函数作为参数*/

package main 

import "fmt"

func main() {
	callback(1,2,add)
}

func add(i,j int) {
	fmt.Printf("%d + %d = %d",i,j,i+j)
}

func callback(i,j int, f func(int, int)) {
	f(i,j)
}
package main 

import (
		"fmt"
		"math/rand"
		//"time"
)

func main() {
	for i:= 0; i < 10; i++ {
		a := rand.Intn(8)	//产生0-8的随机数
		fmt.Printf("%d ",a)
	}
}
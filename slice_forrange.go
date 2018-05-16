package main 

import "fmt"

func main() {
	s1 := make([]int, 5)
	
	for i:= 0; i<len(s1); i++ {
		s1[i] = i*2
	}

	for i,val := range s1 {
		fmt.Printf("%dth is %d\n", i, val)
	}

	s2 := []string{"I", "am", "a", "student"}

	for _,val := range s2 {
		fmt.Println(val)
	}
}
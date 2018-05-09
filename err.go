package main 

import "fmt"

func main() {
	var a int
	var err bool
	a, err = Myfunc(8)
	if err {
		fmt.Println(a)
	} else {
		fmt.Println("error")
	}
	
}


/*先是参数列表，再是返回值列表*/
func Myfunc (a int) (int, bool) {
	if a==0 {
		return -1,false
	}
	return a,true
}
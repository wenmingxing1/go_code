/*通过反射设置值*/

package main 

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Println("settability of v:", v.CanSet())	//判断v是否是可设置的

	v = reflect.ValueOf(&x)		//设成指针
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()	//使用Elem函数将v变为可设置的
	fmt.Println("The Elem of v is:", v)
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(3.1415)	//可以设置了
	fmt.Println(v.Interface())
	fmt.Println(v)
}
/*反射struct*/

package main 

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

//
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)

	fmt.Println(typ)	//NotknownType
	knd := value.Kind()
	fmt.Println(knd)	//struct

	for i:=0; i<value.NumField(); i++ {
		fmt.Printf("Field %d : %v\n", i, value.Field(i))
	}

	//可以使用索引0的方式调用其中的方法
	results := value.Method(0).Call(nil)
	fmt.Println(results)
}
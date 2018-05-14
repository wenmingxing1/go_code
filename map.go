/*map*/

package main 
import "fmt"

func main() {
	m := make(map[string]int)	//使用马克创建一个空的map

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)

	fmt.Println(len(m))

	delete(m, "two")	//删除two

	fmt.Println(m)

	m1 := map[string]int{"one":1, "two":2, "three":3}	//直接创建一个map
	fmt.Println(m)

	for key,val := range m1 {	//利用for-range结构输出key-value
		fmt.Printf("%s == > %d  ", key, val)
	}

}
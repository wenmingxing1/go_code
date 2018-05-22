// 看map中的key是否存在  

package main 
import "fmt"

func main() {
	map1 := make(map[string]int)

	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	value, isPresent := map1["four"]	//isPresent返回一个bool值

	if isPresent {
		fmt.Println(value)
	} else {
		fmt.Println(isPresent)
	}

	
	value, isPresent = map1["two"]

	if isPresent {
		fmt.Println(value)
	} else {
		fmt.Println(isPresent)
	}
	

	delete(map1, "one")		//删除键值对
	fmt.Println(map1)
}
package main 
import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	for key,value := range map1 {
		fmt.Printf("%s is %d\n", key, value)
	}
}
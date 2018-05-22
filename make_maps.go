package main 

import "fmt"

func main() {
	var mapList map[string]int

	var mapAssigned map[string]int

	mapList = map[string]int{"one":1, "two":2}
	mapCreate := make(map[string]float32)
	mapAssigned = mapList

	mapCreate["key1"] = 4.5
	mapCreate["key2"] = 3.14
	mapCreate["key3"] = 3

	fmt.Println(mapList)
	fmt.Println(mapAssigned)
	fmt.Println(mapCreate)
}
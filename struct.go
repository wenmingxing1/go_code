package main 

import "fmt"

type Person struct {
	name string
	age int
	email string
}

func main() {
	person := Person{"Wenmingxing", 24, "ligang1_1@126.com"}
	//var person = Person{name:"Wenmingxing", age:24, email:"ligang1_1@126.com"}

	fmt.Println(person)

	pPerson := &person
	fmt.Println(pPerson)	//&{Wenmingxing 24 ligang1_1@126.com}

	pPerson.age = 50

	fmt.Println(person)
}
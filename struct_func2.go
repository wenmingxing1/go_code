/*结构体方法*/
package main 
import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName string
}

func upPerson(p *Person) {	//可以改变p
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Println(pers1)

	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	upPerson(pers2)
	fmt.Println(pers2)

	pers3 := &Person{"Chris", "Woodward"}	//因为要传入的参数为引用
	upPerson(pers3)
	fmt.Println(pers3)
}
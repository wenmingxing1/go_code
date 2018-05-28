/*由于类型与其方法必须定义在同一个包中
 * 所以只能通过类型别名的方式，在包外定义方法
*/

 package main 
 import (
 	"fmt"
 	"time"
 )

 type myTime struct {
 	time.Time   //anonymous field
 }

 func (t myTime) first3Chars() string {
 	return t.Time.String()[0:3]
 }

 func main() {
 	m := myTime{time.Now()}
 	//调用匿名Time上的String方法
 	fmt.Println("Full time now:", m.String())
 	//调用myTime.first3Chars
 	fmt.Println("First 3 chars:", m.first3Chars())
 }
/*goroutine的使用，如果去掉关键字go
 *则shrotWait会在longWait之后执行
 *而通过go开启goroutine将会并行执行
*/

package main 
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")

	go longWait()	//通过go关键字开启协程
	go shortWait()

	fmt.Println("About to sleep in main()")

	time.Sleep(10*1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5*1e9)
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2*1e9)
	fmt.Println("End of shortWait()")
}


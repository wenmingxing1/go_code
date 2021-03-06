/*缓冲channel
 *
*/

package main 
import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 50)	//buffer = 50
	go func() {
		time.Sleep(15*1e9)
		x := <- c
		fmt.Println("receive", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)

	//time.Sleep(2*1e9)
}
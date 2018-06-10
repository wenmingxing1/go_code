/*channel阻塞测试
 *由于channel没有接收者，所以只会输出一个0，其他的阻塞
*/

package main

import "fmt"
import "time"

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	//go suck(ch1)
	fmt.Println(<-ch1)

	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i:= 0; i <=100 ; i++ {
		ch <- i
	}
}

//suck用于读取ch中的数据，从而不会出现阻塞，也就是说ch中可以持续写入
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
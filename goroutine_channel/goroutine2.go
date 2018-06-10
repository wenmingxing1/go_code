/*channel的使用，完成goroutine的通信*/

package main 

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)	//创建一个string channel ch

	//创建两个goroutine，一个用于发送数据，一个用于接收数据
	go sendDate(ch)

	go getDate(ch)

	time.Sleep(1e9)
}

func sendDate(ch chan string) {
	ch <- "Washington"	// <-表示将值发送到channel中
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getDate(ch chan string) {
	var input string

	//这个for相当于while(1)
	for {
		input = <- ch 		// =<-表示接收来自channel的值
		fmt.Printf("%s ", input)
	}
}
package main 
import "fmt"

func main() {
	var str string = "I am a great man."
	for pos,ch:=range str {
		fmt.Printf("%d is : %c\n", pos, ch)
	}
}
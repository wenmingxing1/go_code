package main 
import "fmt"

func main() {
LABEL:
	for i:=0; i<=5; i++ {
		for j:=0; j<=5; j++ {
			if j==4 {
				continue LABEL
			}
			fmt.Println("i is %d, j is %d",i,j)
		}
	}
}
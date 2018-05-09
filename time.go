package main 

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t:=time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())
}
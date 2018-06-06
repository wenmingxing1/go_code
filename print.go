/*Printf中本来应该使用reflect来推导类型
 * 本例中，我们使用switch来完成
*/

package main 

import (
	"os"
	"strconv"
)

//接口
type Stringer interface {
	String() string 
}

type Celsius float64 

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

type Day int 

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}  

func print(args ...interface{}) {
	for i,arg := range args {
		if i > 0 {os.Stdout.WriteString(" ")}
		switch a := arg.(type) {
			case Stringer: os.Stdout.WriteString(a.String())
			case int: os.Stdout.WriteString(strconv.Itoa(a))
			case string: os.Stdout.WriteString(a)

			default: os.Stdout.WriteString("???")
		}
	}
}

func main() {
	print(Day(1), "was", Celsius(18.36))
}


/*HasPrefix函数用于判断字符串开头*/
/*Contains用于判断字符串中是否包含子串*/
/*Index函数用于返回字符串中子串中第一个字符的索引位置*/
/*LastIndex用于返回子串中第一个字符最后出现的索引位置*/

/*Replace用于将字符串中前n个字符串old替换为new并返回新的字符串*/
/*Count用于统计子串出现的次数*/

/*Fields利用1个或多个空白符号来分割字符串，返回slice*/
/*Split可以指定字符串作为分割的标准*/
/*Join将slice以指定字符拼接*/

package main 

import (
	"fmt"
	"strings"	//string包	
)

func main() {
	var str string = "This is an example of a string"

	fmt.Println(strings.HasPrefix(str, "Th"))
	fmt.Println(strings.Contains(str, " a"))
	fmt.Println(strings.Index(str, "s"))
	fmt.Println(strings.LastIndex(str,"s"))

	fmt.Println(strings.Replace(str, "is", "haha", 24))
	fmt.Println(str)

	fmt.Println(strings.Count(str,"is"))

	fmt.Println(strings.Fields(str))
	fmt.Println(strings.Split(str, "a"))

	fmt.Println(strings.Join(strings.Split(str, "a"), "|"))

}
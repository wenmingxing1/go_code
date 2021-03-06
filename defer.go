/*defer关键字通常用于执行一些函数完成之后的收尾工作
 * 如关闭文件流，打印最终报告等。
*/
package main 
import "fmt"

func main() {
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()	//对比两种情况下的输出结果，defer关键字会将function2放到function1结束之前执行
	//function2()
	fmt.Printf("In function1 at the bottom\n")
}

func function2() {
	fmt.Printf("function2:Deferred until the end of the calling functiong\n")
}
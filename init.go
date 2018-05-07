package trans	//自己构建的包trans

import "math"

var Pi float64


/* init函数不能被人调用，而是在每个包完成初始化后自动执行
 * 一个可能的用途就是在正式开始执行main之前进行数据检测或修复
 * 这个程序不能run，但是可以等待其他main程序调用其中的内容
*/
func init() {
	Pi = 4*math.Atan(1)
}
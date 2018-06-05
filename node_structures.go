/*使用空接口定义节点，作为数据字段的类型*/

package main 
import "fmt"

type Node struct {
	le *Node
	data interface{}
	ri *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

//填充数据字段
func (n *Node) SetData(data interface{}) {
	n.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("root node")

	a := NewNode(nil, nil)
	a.SetData("left node")
	root.le = a

	b := NewNode(nil, nil)
	b.SetData("right node")
	root.ri = b

	fmt.Printf("%v\n",root)
	fmt.Printf("%v\n",root.le)	//a
	fmt.Printf("%v\n",root.ri)	//b

}
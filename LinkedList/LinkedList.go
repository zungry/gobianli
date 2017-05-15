package LinkedList
//http://www.cnblogs.com/requelqi/p/3691711.html

type ElementType interface{}

//IHost is Every node in LinkedList
type INode struct {
	X    ElementType
	next *INode
}

//LinkedList struct
type LinkedList struct {
	Head *INode // head node
}

//func NewIHost
func NewINode(x ElementType, next *INode) *INode {
	return &INode{x, next}
}

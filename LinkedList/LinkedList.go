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

//NewINode return a New Node
func NewINode(x ElementType, next *INode) *INode {
	return &INode{x, next}
}

//NewLinkedList return a New LinkedList
func NewLinkedList() *LinkedList {
	head := &INode{0, nil}
	return &LinkedList{head}
}

//IsEmpty return if a LinkedList is nil
func (list *LinkedList) IsEmpty() bool {
	return list.Head.next == nil
}

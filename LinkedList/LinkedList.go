package LinkedList

import (
	"reflect"
)

//http://www.cnblogs.com/requelqi/p/3691711.html

type ElementType interface{}

//IHost is Every node in LinkedList
type INode struct {
	X    ElementType
	next *INode
}

//LinkedList struct
type LinkedList struct {
	Head *INode // head node, head node store the length of the LinkedList
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

//Length return the length of the LinkedList
func (list *LinkedList) Length() int {
	return int(reflect.ValueOf(list.Head.X).Int()) // use reflect to manipulate objects with arbitary types.
}

//Append append the node at the end of LinkedList
func (list *LinkedList) Append(node *INode) {

}

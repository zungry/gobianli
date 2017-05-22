package LinkedList

import (
	"errors"
	"fmt"
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
	current := list.Head
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	current.next = node
	list.sizeInc()
}

//Prepend nodes
func (list *LinkedList) Prepend(node *INode) {
	current := list.Head
	node.next = current.next
	current.next = node
	list.sizeInc()
}

//Find node
func (list *LinkedList) Find(x ElementType) (*INode, bool) {
	empty := list.IsEmpty()
	if empty {
		fmt.Println("This is an empty list")
		return nil, false
	}
	current := list.Head
	for current.next != nil {
		if current.X == x {
			return current, true
		}
	}
	if current.X == x {
		return current, true
	}
	return nil, false
}

// Remove node
func (list *LinkedList) Remove(x ElementType) error {
	empty := list.IsEmpty()
	if empty {
		return errors.New("This is an empty list")
	}
	current := list.Head
	for current.next != nil {
		if current.next.X == x {
			current.next = current.next.next
			list.sizeDec()
			return nil
		}
		current = current.next
	}
	return nil
}

//sizeInc
func (list *LinkedList) sizeInc() {
	v := int(reflect.ValueOf((*list.Head).X).Int())
	list.Head.X = v + 1
}

//sizeDnc
func (list *LinkedList) sizeDec() {
	v := int(reflect.ValueOf((*list.Head).X).Int())
	list.Head.X = v - 1
}

/**
  打印链表信息
*/
func (list *LinkedList) PrintList() {
	empty := list.IsEmpty()
	if empty {
		fmt.Println("This is an empty list")
		return
	}
	current := list.Head.next
	fmt.Println("The elements is:")
	i := 0
	for ; ; i++ {
		if current.next == nil {
			break
		}
		fmt.Printf("INode%d ,value:%v -> ", i, current.X)
		current = current.next
	}
	fmt.Printf("Node%d value:%v", i+1, current.X)
	return

}

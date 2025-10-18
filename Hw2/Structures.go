package Hw2

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type LinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func (list *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}
	list.size++

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	list.Tail.Next = newNode
	list.Tail = newNode
}

func (list *LinkedList) Insert(index int, value int) {

}

func (list *LinkedList) Remove(index int) {
	if index < 0 || index > list.size {
		/// NOTE(Dima): Error!
		/// TODO(Dima): Maybe return error
		return
	}

	dummyHead := &Node{0, list.Head}
	dummyTail := &Node{0, nil}
	list.Tail.Next = dummyTail

	prev := dummyHead
	curr := list.Head

	for i := 0; i != index; i++ {
		prev = curr
		curr = curr.Next
	}
	prev.Next = curr.Next

	list.size++
	list.Head = dummyHead.Next
	list.Tail.Next = nil
}

func (list *LinkedList) Print() {
	node := list.Head
	for node != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println("nil")
}

func NewLinkedList(vals []int) *LinkedList {
	list := &LinkedList{
		Head: nil,
		Tail: nil,
		size: 0,
	}

	for _, value := range vals {
		list.Append(value)
	}

	return list
}

//type Queue struct {
//	data *LinkedList
//}
//func NewQeue() *Queue {}
//func (qe *Queue) Push(value int) {
//
//}
//
//func (qe *Queue)Pop() {
//
//}

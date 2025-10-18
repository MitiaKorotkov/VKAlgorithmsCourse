package Hw2

import "fmt"

/**
 * Удалить элемент из односвязного списка
 *
 * Необходимо написать функцию, которая
 * принимает на вход односвязный список
 * и некоторое целое число val. Необходимо
 * удалить узел из списка, значение которого
 * равно val.
 */

func RemoveValue(list *LinkedList, value int) {
	dummyHead := &Node{0, list.Head}
	dummyTail := &Node{0, nil}
	list.Tail.Next = dummyTail

	prev := dummyHead
	curr := list.Head

	for curr != nil {
		if curr.Value == value {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	list.size--
	list.Head = dummyHead.Next
	list.Tail.Next = nil
}

func Problem4() {
	list := NewLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	RemoveValue(list, 4)
	RemoveValue(list, 1)
	RemoveValue(list, 9)
	fmt.Print("List: ")
	list.Print()
}

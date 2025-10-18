package Hw2

/**
 * Развернуть односвязный список
 *
 * Необходимо написать функцию, которая
 * принимает на вход односвязный список
 * и разворачивает его.
 */

func ReverseLinkedList(list *LinkedList) *LinkedList {
	if HasCycle(list) {
		return list
	}

	head := list.Head
	cur := list.Head
	var prev *Node = nil

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	list.Head = prev
	list.Tail = head

	return list
}

func Problem2() {
	list := NewLinkedList([]int{1, 2, 3, 4, 5})
	ReverseLinkedList(list)
	list.Print()
}

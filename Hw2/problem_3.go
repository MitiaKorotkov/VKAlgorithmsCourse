package Hw2

import "fmt"

/**
 * Найти середину списка
 *
 * Дан связный список.
 * Необходимо найти середину списка.
 * Сделать это необходимо за O(n)
 * без дополнительных аллокаций
 */

func FindMiddle(list *LinkedList) *Node {
	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func Problem3() {
	list := NewLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	middle := FindMiddle(list)
	fmt.Print("List: ")
	list.Print()
	fmt.Printf("Middle element: %d", middle.Value)
}

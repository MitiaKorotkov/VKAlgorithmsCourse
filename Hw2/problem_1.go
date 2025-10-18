package Hw2

import "fmt"

/**
 * Проверить, является ли список циклическим
 *
 * Дан односвязный список.
 * Необходимо проверить, является ли этот список
 * циклическим.
 *
 * Циклическим (кольцевым) списком называется список
 * у которого последний узел ссылается
 * на один из предыдущих узлов.
 */

func HasCycle(list *LinkedList) bool {
	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

func Problem1() {
	list := NewLinkedList([]int{1, 2})
	fmt.Println(HasCycle(list))
}

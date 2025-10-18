package Hw2

/**
 * Слияние двух отсортированных списков
 *
 * Написать функцию, которая принимает на вход
 * два отсортированных односвязных списка и
 * объединяет их в один отсортированный список.
 * При этом затраты по памяти должны быть O(1)
 */

func mergeTwoSortedLists(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	dummyHead := &Node{0, nil}
	tmp := dummyHead

	first := l1.Head
	second := l2.Head

	for ; first != nil && second != nil; tmp = tmp.Next {
		if first.Value <= second.Value {
			tmp.Next = first
			first = first.Next
		} else {
			tmp.Next = second
			second = second.Next
		}
	}

	if first != nil {
		tmp.Next = first
	}
	if second != nil {
		tmp.Next = second
	}
	tmp = tmp.Next

	return &LinkedList{Head: dummyHead.Next, Tail: tmp, size: l1.size + l2.size}
}

func Problem7() {
	l1 := NewLinkedList([]int{1, 3, 5, 6, 7})
	l2 := NewLinkedList([]int{1, 2, 3})

	mergeTwoSortedLists(l1, l2)
	l1.Print()
}

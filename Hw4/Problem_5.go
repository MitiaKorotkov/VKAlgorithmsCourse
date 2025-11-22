package Hw4

import "fmt"

/*
 * Сравнение двух деревьев
 *
 * На вход функции подается 2 бинарных дерева. Необходимо понять,
 * являются ли эти два дерева одинаковыми.
 */

func isSameTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Data != b.Data {
		return false
	}

	return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}

func Problem5() {
	arr1 := []int{8, 9, 11, 7, 16, 3, 1}
	arr2 := []int{8, 9, 11, 7, 16, 3, 1}
	tree1 := buildTree(arr1, 0)
	tree2 := buildTree(arr2, 0)

	if isSameTree(tree1, tree2) {
		fmt.Print("Trees are same")
	} else {
		fmt.Print("Trees not same")
	}
}

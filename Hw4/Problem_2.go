package Hw4

import "fmt"

/*
 * Симметричное бинарное дерево
 *
 * На вход функции подается бинарное дерево. Необходимо понять,
 * является ли это дерево симметричным.
 */

func symmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return symmetricHelper(left.Left, right.Right) &&
		symmetricHelper(left.Right, right.Left) &&
		left.Data == right.Data
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmetricHelper(root.Left, root.Right)
}

func Problem2() {
	arr := []int{8, 9, 11, 7, 16, 3, 1}
	tree := buildTree(arr, 0)

	if isSymmetric(tree) {
		fmt.Println("Tree is symmetric")
	} else {
		fmt.Println("Tree is not symmetric")
	}
}

package Hw4

import "fmt"

/*
 * Восстановление бинарного дерева из массива
 *
 * Необходимо реализовать функцию, которая будет принимать
 * на вход массив и выстраивать из него бинарное дерево.
 */

func buildTree(arr []int, i int) *TreeNode {
	if i >= len(arr) {
		return nil
	}

	root := TreeNode{arr[i], nil, nil}
	root.Left = buildTree(arr, 2*i+1)
	root.Right = buildTree(arr, 2*i+2)
	return &root
}

func Problem1() {
	arr := []int{8, 9, 11, 7, 16, 3, 1}

	tree := buildTree(arr, 0)
	matrix := printTree(tree)

	for _, row := range matrix {
		fmt.Println(row)
	}
}

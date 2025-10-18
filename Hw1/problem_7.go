package Hw1

import "fmt"

/**
 * Передвинуть четные числа вперед
 *
 * Дан не отсортированный массив целых чисел
 *
 * Необходимо перенести в начало массива все
 * четные числа, сохраняя их очередность
 */

func EvenFirst(arr []int) {
	insertion := 0

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i], arr[insertion] = arr[insertion], arr[i]
			insertion++
		}
	}
}

func Task7() {
	var size int
	fmt.Print("Enter array size: ")
	fmt.Scanln(&size)

	fmt.Print("Enter array: ")
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	EvenFirst(arr)

	fmt.Print("Sorted array: ")
	fmt.Println(arr)
}

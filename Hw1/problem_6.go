package Hw1

/**
 * Задача флага Нидерландов
 *
 * Дан массив состоящий из нулей,
 * единиц и двоек
 *
 * Необходимо его отсортировать
 * за линейное время
 */

import "fmt"

func SortColors(arr []int) {
	left := 0
	mid := 0
	right := len(arr) - 1

	for mid <= right {
		switch arr[mid] {
		case 0:
			arr[mid], arr[left] = arr[left], arr[mid]
			left++
			mid++
		case 1:
			mid++
		case 2:
			arr[mid], arr[right] = arr[right], arr[mid]
			right--
		}
	}
}

func Problem6() {
	var size int
	fmt.Print("Enter array size: ")
	fmt.Scanln(&size)

	fmt.Print("Enter array: ")
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	SortColors(arr)

	fmt.Print("Sorted array: ")
	fmt.Println(arr)
}
